package util

type HttpMethod string
type HttpStatusCode int
type HttpVersion string
type HttpMediaType string
type HttpHeader string
type HttpConnection string
type HttpTransferEncoding string
type HttpExpect string

const (
	MethodGet     HttpMethod = "GET"
	MethodHead    HttpMethod = "HEAD"
	MethodPost    HttpMethod = "POST"
	MethodPut     HttpMethod = "PUT"
	MethodDelete  HttpMethod = "DELETE"
	MethodConnect HttpMethod = "CONNECT"
	MethodOptions HttpMethod = "OPTIONS"
	MethodTrace   HttpMethod = "TRACE"
)

const (
	StatusContinue HttpStatusCode = iota + 100
	StatusSwitchingProtocols
	StatusProcessing
)

const (
	StatusOK HttpStatusCode = iota + 200
	StatusCreated
	StatusAccepted
	StatusNonAuthoritativeInformation
	StatusNoContent
	StatusResetContent
	StatusPartialContent
	StatusMultiStatus
	StatusAlreadyReported
	StatusIMUsed
)

const (
	StatusMultipleChoices HttpStatusCode = iota + 300
	StatusMovedPermanently
	StatusFound
	StatusSeeOther
	StatusNotModified
	StatusUseProxy
	StatusTemporaryRedirect
	StatusPermanentRedirect
)

const (
	StatusBadRequest HttpStatusCode = iota + 400
	StatusUnauthorized
	StatusPaymentRequired
	StatusForbidden
	StatusNotFound
	StatusMethodNotAllowed
	StatusNotAcceptable
	StatusProxyAuthenticationRequired
	StatusRequestTimeout
	StatusConflict
	StatusGone
	StatusLengthRequired
	StatusPreconditionFailed
	StatusEntityTooLarge
	StatusRequestURITooLong
	StatusUnsupportedMediaType
	StatusRequestedRangeNotSatisfiable
	StatusExpectationFailed
	StatusImATeapot
	StatusMisdirectedRequest              = 421
	StatusUnprocessableEntity             = 422
	StatusLocked                          = 423
	StatusFailedDependency                = 424
	StatusUpgradeRequired                 = 426
	StatusPreconditionRequired            = 428
	StatusTooManyRequests                 = 429
	StatusRequestHeaderFieldsTooLarge     = 431
	StatusConnectionClosedWithoutResponse = 444
	StatusUnavailableForLegalReasons      = 451
	StatusClientClosedRequest             = 499
)

const (
	StatusInternalServerError HttpStatusCode = iota + 500
	StatusNotImplemented
	StatusBadGateway
	StatusServiceUnavailable
	StatusGatewayTimeout
	StatusHTTPVersionNotSupported
	StatusVariantAlsoNegotiates
	StatusInsufficientStorage
	StatusLoopDetected
	StatusNotExtended                   = 510
	StatusNetworkAuthenticationRequired = 511
	StatusNetworkConnectTimeoutError    = 599
)

const (
	Version09 HttpVersion = "HTTP/0.9"
	Version10 HttpVersion = "HTTP/1.0"
	Version11 HttpVersion = "HTTP/1.1"
)

const (
	MediaTypeAAC        HttpMediaType = "audio/aac"
	MediaTypeAVI        HttpMediaType = "video/x-msvideo"
	MediaTypeBinary     HttpMediaType = "application/octet-stream"
	MediaTypeBitmap     HttpMediaType = "image/bmp"
	MediaTypeCSS        HttpMediaType = "text/css"
	MediaTypeCSV        HttpMediaType = "text/csv"
	MediaTypeEPUB       HttpMediaType = "application/epub+zip"
	MediaTypeGZip       HttpMediaType = "application/gzip"
	MediaTypeGIF        HttpMediaType = "image/gif"
	MediaTypeHTML       HttpMediaType = "text/html"
	MediaTypeHTTP       HttpMediaType = "message/http"
	MediaTypeIcon       HttpMediaType = "image/vnd.microsoft.icon"
	MediaTypeJPEG       HttpMediaType = "image/jpeg"
	MediaTypeJavaScript HttpMediaType = "text/javascript"
	MediaTypeJSON       HttpMediaType = "application/json"
	MediaTypeMP3        HttpMediaType = "audio/mpeg"
	MediaTypeMP4        HttpMediaType = "video/mp4"
	MediaTypeOGGAudio   HttpMediaType = "audio/ogg"
	MediaTypePNG        HttpMediaType = "image/png"
	MediaTypePDF        HttpMediaType = "application/pdf"
	MediaTypePHP        HttpMediaType = "application/php"
	MediaTypeRTF        HttpMediaType = "application/rtf"
	MediaTypeSVG        HttpMediaType = "image/svg+xml"
	MediaTypeSWF        HttpMediaType = "application/x-shockwave-flash"
	MediaTypeTTF        HttpMediaType = "font/ttf"
	MediaTypeText       HttpMediaType = "text/plain"
	MediaTypeWAV        HttpMediaType = "audio/wav"
	MediaTypeWEBMAudio  HttpMediaType = "audio/webm"
	MediaTypeWEBMVideo  HttpMediaType = "video/webm"
	MediaTypeWEBPImage  HttpMediaType = "image/webp"
	MediaTypeWOFF       HttpMediaType = "font/woff"
	MediaTypeWOFF2      HttpMediaType = "font/woff2"
	MediaTypeXHTML      HttpMediaType = "application/xhtml+xml"
	MediaTypeXML        HttpMediaType = "application/xml"
	MediaTypeZip        HttpMediaType = "application/zip"
)

const (
	HeaderHost             HttpHeader = "host"
	HeaderConnection       HttpHeader = "connection"
	HeaderContentLength    HttpHeader = "content-length"
	HeaderContentType      HttpHeader = "content-type"
	HeaderTransferEncoding HttpHeader = "transfer-encoding"
	HeaderTE               HttpHeader = "te"
	HeaderUpgrade          HttpHeader = "upgrade"
	HeaderVia              HttpHeader = "via"
	HeaderLastModified     HttpHeader = "last-modified"
	HeaderETag             HttpHeader = "etag"
	HeaderExpect           HttpHeader = "expect"
	HeaderServer           HttpHeader = "server"
)

const (
	ConnectionClose     HttpConnection = "close"
	ConnectionKeepAlive HttpConnection = "keep-alive"
)

const (
	TransferEncodingChunked  HttpTransferEncoding = "chunked"
	TransferEncodingCompress HttpTransferEncoding = "compress"
	TransferEncodingIdentity HttpTransferEncoding = "identity"
	TransferEncodingDeflate  HttpTransferEncoding = "deflate"
	TransferEncodingGZip     HttpTransferEncoding = "gzip"
)

const HttpExpectContinue HttpExpect = "100-continue"

type UriForm int
type UriScheme string

const (
	UriFormOrigin UriForm = iota
	UriFormAbsolute
	UriFormAuthority
	UriFormAsterisk
)

const (
	UriSchemeHttp  UriScheme = "http"
	UriSchemeHttps UriScheme = "https"
)
