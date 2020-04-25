package response

import (
	"bufio"
	"fmt"
	"log"
	"segaline/src/util"
	"strconv"
)

type Response struct {
	HttpVersion util.HttpVersion
	StatusCode  util.HttpStatusCode

	Headers map[util.HttpHeader]string
	Body    []byte
	Chunked bool
}

func New() *Response {
	return &Response{
		HttpVersion: util.DefaultHttpVersion,
		Headers: map[util.HttpHeader]string{
			util.HeaderContentLength: "0",
			util.HeaderServer:        util.ServerNameVersion,
		},
	}
}

func (res *Response) WithStatus(status util.HttpStatusCode) *Response {
	res.StatusCode = status
	if int(status) < 200 || int(status) == 204 {
		delete(res.Headers, util.HeaderContentLength)
	}
	return res
}

func (res *Response) WithHeader(header util.HttpHeader, value string) *Response {
	res.Headers[header] = value
	return res
}

func (res *Response) WithoutHeader(header util.HttpHeader) *Response {
	delete(res.Headers, header)
	return res
}

func (res *Response) WithBody(body []byte, mediaType util.HttpMediaType) *Response {
	res.Body = body
	res.WithHeader(util.HeaderContentType, string(mediaType))

	if len(body) > util.ResponseMaxUnchunkedBody {
		res.Chunked = true
		return res.
			WithoutHeader(util.HeaderContentLength).
			WithHeader(util.HeaderTransferEncoding, string(util.TransferEncodingChunked))
	} else {
		return res.WithHeader(util.HeaderContentLength, strconv.Itoa(len(body)))
	}
}

func (res *Response) AsBytesWithoutBody() []byte {
	headers := ""
	for name, value := range res.Headers {
		headers += string(name) + ": " + value + "\r\n"
	}

	str := fmt.Sprintf("%s %d\r\n%s\r\n", res.HttpVersion, res.StatusCode, headers)
	return []byte(str)
}

func (res *Response) AsBytes() []byte {
	str := fmt.Sprintf("%s%s", res.AsBytesWithoutBody(), res.Body)
	return []byte(str)
}

func (res *Response) Respond(writer *bufio.Writer) {
	if res.Chunked {
		writeFullyLog(writer, res.AsBytesWithoutBody())

		chunkSize := util.ResponseChunkSize
		written := 0
		for written < len(res.Body)/chunkSize*chunkSize {
			buf := []byte(fmt.Sprintf("%x\r\n%s\r\n", chunkSize, res.Body[written:written+chunkSize]))
			writeFullyLog(writer, buf)
			flushLog(writer)
			written += chunkSize
		}

		buf := []byte(fmt.Sprintf("%x\r\n%s\r\n0\r\n\r\n", len(res.Body)%chunkSize, res.Body[written:]))
		writeFullyLog(writer, buf)
		flushLog(writer)
	} else {
		writeFullyLog(writer, res.AsBytes())
		flushLog(writer)
	}
}

func RespondStatus(writer *bufio.Writer, status util.HttpStatusCode, closeConnection bool) {
	res := New().WithStatus(status)
	if closeConnection {
		res.WithHeader(util.HeaderConnection, string(util.ConnectionClose))
	}
	res.Respond(writer)
}

func writeFullyLog(writer *bufio.Writer, bytes []byte) int {
	written, err := writeFully(writer, bytes)
	if err != nil {
		log.Println("An issue occurred while responding to a request.")
	}
	return written
}

func writeFully(writer *bufio.Writer, bytes []byte) (int, error) {
	written := 0
	for written < len(bytes) {
		n, err := writer.Write(bytes[written:])
		if err != nil {
			return written, err
		}
		written += n
	}
	return written, nil
}

func flushLog(writer *bufio.Writer) {
	if err := writer.Flush(); err != nil {
		log.Println("An issue occurred while responding to a request.")
	}
}
