package util

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
	DefaultEmptyRequestTarget = "/index.html"
	DefaultHttpVersion        = HttpVersion11
)

const (
	HeaderHost             = "Host"
	HeaderConnection       = "Connection"
	HeaderContentLength    = "Content-Length"
	HeaderContentType      = "Content-Type"
	HeaderTransferEncoding = "Transfer-Encoding"
)

const (
	ErrorContentLengthExceeded    = "content length maximum exceeded"
	ErrorRequestURILengthExceeded = "request uri length maximum exceeded"
)
