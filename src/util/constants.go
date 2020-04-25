package util

import "time"

const (
	DefaultEmptyRequestTarget = "/index.html"
	DefaultHttpVersion        = HttpVersion11
	DefaultReadTimeout        = 10 * time.Second
)

const (
	RequestMaxContentLength = 65_536
	RequestMaxURILength     = 32_768
)

const (
	ResponseWriterBufferSize = 4_096
	ResponseChunkSize        = 4_096
	ResponseMaxUnchunkedBody = 8 * ResponseChunkSize
)

const (
	HeaderHost             = "host"
	HeaderConnection       = "connection"
	HeaderContentLength    = "content-length"
	HeaderContentType      = "content-type"
	HeaderTransferEncoding = "transfer-encoding"
	HeaderExpect           = "expect"
)

const (
	ErrorContentLengthExceeded       = "content length maximum exceeded"
	ErrorRequestURILengthExceeded    = "request uri length maximum exceeded"
	ErrorUnsupportedTransferEncoding = "unsupported transfer encoding"
	ErrorTimeoutReached              = "timeout reached"
)
