package http

import (
	"bufio"
	"fmt"
	"log"
	"segaline/src/util"
	"strconv"
	"time"
)

type Response struct {
	HttpVersion Version
	StatusCode  StatusCode

	Headers map[Header]string
	Body    []byte
	Chunked bool

	request *Request
}

func NewResponse(req *Request) *Response {
	return &Response{
		HttpVersion: Version11,
		Headers: map[Header]string{
			HeaderContentLength: "0",
			HeaderServer:        util.ServerNameVersion,
			HeaderDate:          formatTimeGMT(time.Now()),
		},
		request: req,
	}
}

func (res *Response) WithStatus(status StatusCode) *Response {
	res.StatusCode = status
	if int(status) < 200 || int(status) == 204 {
		res.WithoutHeader(HeaderContentLength)
	}
	return res
}

func (res *Response) WithHeader(header Header, value string) *Response {
	res.Headers[header] = value
	return res
}

func (res *Response) WithoutHeader(header Header) *Response {
	delete(res.Headers, header)
	return res
}

func (res *Response) WithBody(body []byte, mediaType MediaType) *Response {
	res.Body = body
	res.WithHeader(HeaderContentType, string(mediaType))

	if len(body) > util.ResponseMaxUnchunkedBody {
		res.Chunked = true
		return res.
			WithoutHeader(HeaderContentLength).
			WithHeader(HeaderTransferEncoding, string(TransferEncodingHeaderChunked))
	} else {
		return res.WithHeader(HeaderContentLength, strconv.Itoa(len(body)))
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

	if res.StatusCode != StatusRequestTimeout {
		log.Printf("(%d) %s %s %s\n", res.StatusCode, res.request.Method, &res.request.Uri, res.request.RemoteAddr)
	}
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
