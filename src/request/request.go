package request

import (
	"bufio"
	"io"
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

func Parse(reader io.Reader) (Request, error) {
	parser := newRequestParser(bufio.NewReader(reader))
	return parser.parse()
}
