package http

import (
	"bufio"
	"fmt"
	"net"
)

type Request struct {
	Method      Method
	Uri         Uri
	HttpVersion Version

	Headers map[string]string
	Body    []byte

	RemoteAddr net.Addr
}

func ParseRequest(conn net.Conn) (Request, error) {
	parser := newRequestParser(bufio.NewReader(conn), bufio.NewWriter(conn))
	return parser.parse(conn.RemoteAddr())
}

func (req *Request) WillCloseConnection() bool {
	value, ok := req.Headers[string(HeaderConnection)]
	hasClose := ok && value == string(ConnectionHeaderClose)
	isKeepAlive := req.HttpVersion == Version10 && ok && value == string(ConnectionHeaderKeepAlive)
	return hasClose || req.HttpVersion < Version11 && !isKeepAlive
}

func (req *Request) AsBytes() []byte {
	headers := ""
	for name, value := range req.Headers {
		headers += name + ": " + value + "\r\n"
	}

	str := fmt.Sprintf("%s %s %s\r\n%s\r\n%s", req.Method, &req.Uri, req.HttpVersion, headers, req.Body)
	return []byte(str)
}
