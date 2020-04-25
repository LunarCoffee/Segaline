package util

import "time"

const (
	ServerName        = "Segaline"
	ServerVersion     = "0.1.0"
	ServerNameVersion = ServerName + "/" + ServerVersion
)

const (
	DefaultEmptyRequestTarget    = "/index.html"
	DefaultHttpVersion           = Version11
	DefaultReadTimeout           = 10 * time.Second
	DefaultFallbackErrorTemplate = "{statusCode} - {serverInfo}"
)

const (
	RequestMaxContentLength = 65_536
	RequestMaxURILength     = 32_768
)

const (
	ResponseWriterBufferSize = 4_096
	ResponseChunkSize        = 4_096
	ResponseMaxUnchunkedBody = 8 * ResponseChunkSize
	ResponseTimeFormat       = "Mon, 02 Jan 2006 15:04:05"
)

const (
	ErrorContentLengthExceeded       = "content length maximum exceeded"
	ErrorRequestURILengthExceeded    = "request uri length maximum exceeded"
	ErrorUnsupportedTransferEncoding = "unsupported transfer encoding"
	ErrorTimeoutReached              = "timeout reached"
)
