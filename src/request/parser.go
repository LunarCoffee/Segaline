package request

import (
	"bufio"
	"errors"
	"segaline/src/response"
	"segaline/src/uri"
	"segaline/src/util"
	"strconv"
	"strings"
	"time"
)

const optionalWhiteSpace = " \t"

type requestParser struct {
	reader *bufio.Reader
	writer *bufio.Writer
}

func newRequestParser(reader *bufio.Reader, writer *bufio.Writer) requestParser {
	return requestParser{reader, writer}
}

func (parser *requestParser) parse() (request Request, err error) {
	method, reqUri, httpVersion, err := parser.parseRequestLine()
	if err != nil {
		return
	}

	headers, err := parser.parseHeaders()
	if err != nil {
		return
	}
	if _, ok := headers[util.HeaderHost]; !ok {
		err = errors.New("missing host header")
	}

	// Trailer headers are checked for duplication but are ultimately ignored (as with most implementations).
	body, trailer, err := parser.parseBody(headers)
	if err != nil {
		return
	}
	for key := range trailer {
		if _, ok := headers[key]; ok {
			err = errors.New("duplicate header in trailer")
			return
		}
	}

	return Request{method, reqUri, httpVersion, headers, body}, nil
}

func (parser *requestParser) parseRequestLine() (m util.HttpMethod, u uri.Uri, v util.HttpVersion, err error) {
	line, err := parser.readLine()
	if err != nil {
		return
	}

	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return
	}

	m = util.HttpMethod(parts[0])
	switch m {
	case util.HttpMethodGet, util.HttpMethodHead, util.HttpMethodPost, util.HttpMethodPut, util.HttpMethodDelete,
		util.HttpMethodConnect, util.HttpMethodOptions, util.HttpMethodTrace:
	default:
		err = errors.New("invalid method")
		return
	}

	u, err = uri.Parse(m, parts[1])
	if err != nil {
		return
	}

	v = util.HttpVersion(parts[2])
	switch v {
	case util.HttpVersion09, util.HttpVersion10, util.HttpVersion11:
	default:
		err = errors.New("unsupported http version")
		return
	}
	return
}

func (parser *requestParser) parseHeaders() (headers map[string]string, err error) {
	var line string
	headers = map[string]string{}

	for line, err = parser.readLine(); line != ""; line, err = parser.readLine() {
		if err != nil {
			return
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) < 2 {
			err = errors.New("invalid header")
			return
		}

		name := util.NormalizeCase(parts[0])
		value := strings.Trim(util.NormalizeCase(parts[1]), optionalWhiteSpace)
		if !util.IsVisibleString(name) || !util.IsValidHeaderValue(value) {
			err = errors.New("invalid header")
			return
		}

		if _, ok := headers[name]; ok {
			err = errors.New("duplicate header")
			return
		}
		headers[name] = value
	}

	parser.sendContinue(headers)
	return
}

func (parser *requestParser) parseBody(headers map[string]string) (body []byte, trailer map[string]string, err error) {
	trailer = map[string]string{}

	if rawEncodings, ok := headers[util.HeaderTransferEncoding]; ok {
		if rawEncodings != string(util.HttpTransferEncodingChunked) {
			err = errors.New(util.ErrorUnsupportedTransferEncoding)
			return
		}
		body, trailer, err = parser.readChunked()
	} else if contentLength, ok := headers[util.HeaderContentLength]; ok {
		var length int
		length, err = strconv.Atoi(contentLength)
		if err != nil {
			return
		}
		if length > util.RequestMaxContentLength {
			err = errors.New(util.ErrorContentLengthExceeded)
			return
		}
		body, err = parser.readBytesFully(length)
	}
	return
}

func (parser *requestParser) readChunked() (body []byte, trailer map[string]string, err error) {
	var chunkHeader string
	chunkSize := int64(-1)

	for chunkSize != 0 {
		chunkHeader, err = parser.readLine()
		if err != nil {
			return
		}

		parts := strings.Split(chunkHeader, ";")
		chunkSize, err = strconv.ParseInt(parts[0], 16, 32)
		if err != nil || chunkSize > util.ResponseChunkSize {
			err = errors.New("invalid chunk size")
			return
		}

		if chunkSize > 0 {
			var bytes []byte
			bytes, err = parser.readBytesFully(int(chunkSize))
			if line, readErr := parser.readLine(); readErr != nil || line != "" {
				if readErr == nil {
					err = errors.New("invalid chunk")
				}
				return
			}
			body = append(body, bytes...)
		}
	}

	trailer, err = parser.parseHeaders()
	return
}

func (parser *requestParser) readBytesFully(bytesToRead int) ([]byte, error) {
	var body []byte
	totalRead := 0

	for totalRead < bytesToRead {
		bytes, _, err := parser.rawReadTimeout(
			func() ([]byte, bool, error) {
				b, err := parser.reader.ReadByte()
				return []byte{b}, false, err
			},
		)
		if err != nil {
			return nil, err
		}
		body = append(body, bytes[0])
		totalRead++
	}
	return body, nil
}

func (parser *requestParser) readLine() (string, error) {
	line, isPrefix, err := parser.rawReadTimeout(parser.reader.ReadLine)
	if err != nil {
		return "", err
	}
	fullLine := make([]byte, len(line))
	copy(fullLine, line)

	for isPrefix {
		fragment, curIsPrefix, err := parser.rawReadTimeout(parser.reader.ReadLine)
		if err != nil {
			return "", err
		}
		fullLine = append(fullLine, fragment...)
		isPrefix = curIsPrefix
	}
	return string(fullLine), nil
}

func (parser *requestParser) rawReadTimeout(f func() ([]byte, bool, error)) (line []byte, prefix bool, err error) {
	lineChan := make(chan []byte)
	prefixChan := make(chan bool)
	errChan := make(chan error)

	go func() {
		// If the timeout is reached, the connection will be terminated and this call will fail, so this goroutine
		// won't be blocked indefinitely.
		line, isPrefix, err := f()
		if err != nil {
			errChan <- err
		} else {
			lineChan <- line
			prefixChan <- isPrefix
		}
	}()

	select {
	case line = <-lineChan:
		prefix = <-prefixChan
	case err = <-errChan:
	case <-time.After(util.DefaultReadTimeout):
		err = errors.New(util.ErrorTimeoutReached)
	}
	return
}

func (parser *requestParser) sendContinue(headers map[string]string) {
	if value := headers[util.HeaderExpect]; value == string(util.HttpExpectContinue) {
		response.RespondStatus(parser.writer, util.HttpStatusContinue, false)
	}
}
