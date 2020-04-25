package request

import (
	"bufio"
	"fmt"
	"net"
	"segaline/src/uri"
	"segaline/src/util"
)

type Request struct {
	Method      util.HttpMethod
	Uri         uri.Uri
	HttpVersion util.HttpVersion

	Headers map[string]string
	Body    []byte
}

func Parse(conn net.Conn) (Request, error) {
	parser := newRequestParser(bufio.NewReader(conn), bufio.NewWriter(conn))
	return parser.parse()
}

func (req *Request) WillCloseConnection() bool {
	value, ok := req.Headers[string(util.HeaderConnection)]
	hasClose := ok && value == string(util.ConnectionClose)
	isKeepAlive := req.HttpVersion == util.Version10 && ok && value == string(util.ConnectionKeepAlive)
	return hasClose || req.HttpVersion < util.Version11 && !isKeepAlive
}

func (req *Request) AsBytes() []byte {
	headers := ""
	for name, value := range req.Headers {
		headers += name + ": " + value + "\r\n"
	}

	str := fmt.Sprintf("%s %s %s\r\n%s\r\n%s", req.Method, &req.Uri, req.HttpVersion, headers, req.Body)
	return []byte(str)
}
