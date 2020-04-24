package request

import (
	"bufio"
	"errors"
	"segaline/src/uri"
	"segaline/src/util"
	"strconv"
	"strings"
)

const optionalWhiteSpace = " \t"

type requestParser struct {
	reader *bufio.Reader
}

func newRequestParser(reader *bufio.Reader) requestParser {
	return requestParser{reader}
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
	body, err := parser.parseBody(headers)
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

		name := parts[0]
		value := strings.Trim(parts[1], optionalWhiteSpace)
		if !util.IsVisibleString(name) || !util.IsValidHeaderValue(value) {
			err = errors.New("invalid header")
			return
		}

		_, ok := headers[name]
		if ok {
			err = errors.New("duplicate header")
			return
		}
		headers[name] = value
	}

	if _, ok := headers[util.HeaderHost]; !ok {
		err = errors.New("missing host header")
	}
	return
}

func (parser *requestParser) parseBody(headers map[string]string) (body []byte, err error) {
	if _, ok := headers[util.HeaderTransferEncoding]; ok {
		// TODO:
	} else if contentLength, ok := headers[util.HeaderContentLength]; ok {
		contentLength, err := strconv.Atoi(contentLength)
		if err != nil {
			return nil, err
		}
		if contentLength > util.RequestMaxContentLength {
			return nil, errors.New(util.ErrorContentLengthExceeded)
		}

		totalRead := 0
		for totalRead < contentLength {
			b, err := parser.reader.ReadByte()
			if err != nil {
				return nil, err
			}
			body = append(body, b)
			totalRead++
		}
	}
	return
}

func (parser *requestParser) readLine() (string, error) {
	line, isPrefix, err := parser.reader.ReadLine()
	if err != nil {
		return "", err
	}
	fullLine := make([]byte, len(line))
	copy(fullLine, line)

	for isPrefix {
		fragment, curIsPrefix, err := parser.reader.ReadLine()
		if err != nil {
			return "", err
		}
		fullLine = append(fullLine, fragment...)
		isPrefix = curIsPrefix
	}
	return string(fullLine), nil
}
