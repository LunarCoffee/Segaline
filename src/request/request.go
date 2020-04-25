package request

import (
	"bufio"
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
	value, ok := req.Headers[util.HeaderConnection]
	hasClose := ok && value == string(util.HttpConnectionClose)
	isKeepAlive := req.HttpVersion == util.HttpVersion10 && ok && value == string(util.HttpConnectionKeepAlive)
	return hasClose || req.HttpVersion < util.HttpVersion11 && !isKeepAlive
}
