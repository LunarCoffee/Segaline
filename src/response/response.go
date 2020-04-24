package response

import (
	"fmt"
	"segaline/src/util"
	"strconv"
)

type Response struct {
	HttpVersion util.HttpVersion
	StatusCode  util.HttpStatusCode

	Headers map[string]string
	Body    []byte
}

func New() *Response {
	return &Response{
		HttpVersion: util.HighestSupportedHttpVersion,
		Headers:     map[string]string{},
	}
}

func (res *Response) WithStatus(status util.HttpStatusCode) *Response {
	res.StatusCode = status
	return res
}

func (res *Response) WithHeader(name string, value string) *Response {
	// TODO: percent encode these first, then you dont need the same check
	if !util.IsVisibleString(name) || !util.IsValidHeaderValue(value) {
		panic("invalid header")
	}
	res.Headers[name] = value
	return res
}

func (res *Response) WithBody(body []byte, mediaType util.HttpMediaType) *Response {
	res.Body = body
	return res.
		WithHeader(util.ContentTypeHeader, string(mediaType)).
		WithHeader(util.ContentLengthHeader, strconv.Itoa(len(body)))
}

func (res *Response) AsBytes() []byte {
	var httpVersion, headers string

	switch res.HttpVersion {
	case util.HttpVersion09:
		httpVersion = "HTTP/0.9"
	case util.HttpVersion10:
		httpVersion = "HTTP/1.0"
	case util.HttpVersion11:
		httpVersion = "HTTP/1.1"
	case util.HttpVersion20:
		httpVersion = "HTTP/2.0"
	}

	for name, value := range res.Headers {
		headers += name + ": " + value + "\r\n"
	}

	str := fmt.Sprintf("%s %d\r\n%s\r\n%s", httpVersion, res.StatusCode, headers, res.Body)
	return []byte(str)
}
