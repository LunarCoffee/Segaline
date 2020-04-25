package util

type HttpMethod string
type HttpStatusCode int
type HttpVersion string
type HttpMediaType string
type HttpConnection string
type HttpTransferEncoding string
type HttpExpect string

const (
	HttpMethodGet     HttpMethod = "GET"
	HttpMethodHead    HttpMethod = "HEAD"
	HttpMethodPost    HttpMethod = "POST"
	HttpMethodPut     HttpMethod = "PUT"
	HttpMethodDelete  HttpMethod = "DELETE"
	HttpMethodConnect HttpMethod = "CONNECT"
	HttpMethodOptions HttpMethod = "OPTIONS"
	HttpMethodTrace   HttpMethod = "TRACE"
)

const (
	HttpStatusContinue HttpStatusCode = iota + 100
	HttpStatusSwitchingProtocols
	HttpStatusProcessing
)

const (
	HttpStatusOK HttpStatusCode = iota + 200
	HttpStatusCreated
	HttpStatusAccepted
	HttpStatusNonAuthoritativeInformation
	HttpStatusNoContent
	HttpStatusResetContent
	HttpStatusPartialContent
	HttpStatusMultiStatus
	HttpStatusAlreadyReported
	HttpStatusIMUsed
)

const (
	HttpStatusMultipleChoices HttpStatusCode = iota + 300
	HttpStatusMovedPermanently
	HttpStatusFound
	HttpStatusSeeOther
	HttpStatusNotModified
	HttpStatusUseProxy
	HttpStatusTemporaryRedirect
	HttpStatusPermanentRedirect
)

const (
	HttpStatusBadRequest HttpStatusCode = iota + 400
	HttpStatusUnauthorized
	HttpStatusPaymentRequired
	HttpStatusForbidden
	HttpStatusNotFound
	HttpStatusMethodNotAllowed
	HttpStatusNotAcceptable
	HttpStatusProxyAuthenticationRequired
	HttpStatusRequestTimeout
	HttpStatusConflict
	HttpStatusGone
	HttpStatusLengthRequired
	HttpStatusPreconditionFailed
	HttpStatusPayloadTooLarge
	HttpStatusRequestURITooLong
	HttpStatusUnsupportedMediaType
	HttpStatusRequestedRangeNotSatisfiable
	HttpStatusExpectationFailed
	HttpStatusImATeapot
	HttpStatusMisdirectedRequest              = 421
	HttpStatusUnprocessableEntity             = 422
	HttpStatusLocked                          = 423
	HttpStatusFailedDependency                = 424
	HttpStatusUpgradeRequired                 = 426
	HttpStatusPreconditionRequired            = 428
	HttpStatusTooManyRequests                 = 429
	HttpStatusRequestHeaderFieldsTooLarge     = 431
	HttpStatusConnectionClosedWithoutResponse = 444
	HttpStatusUnavailableForLegalReasons      = 451
	HttpStatusClientClosedRequest             = 499
)

const (
	HttpStatusInternalServerError HttpStatusCode = iota + 500
	HttpStatusNotImplemented
	HttpStatusBadGateway
	HttpStatusServiceUnavailable
	HttpStatusGatewayTimeout
	HttpStatusHTTPVersionNotSupported
	HttpStatusVariantAlsoNegotiates
	HttpStatusInsufficientStorage
	HttpStatusLoopDetected
	HttpStatusNotExtended                   = 510
	HttpStatusNetworkAuthenticationRequired = 511
	HttpStatusNetworkConnectTimeoutError    = 599
)

const (
	HttpVersion09 HttpVersion = "HTTP/0.9"
	HttpVersion10 HttpVersion = "HTTP/1.0"
	HttpVersion11 HttpVersion = "HTTP/1.1"
)

const (
	HttpMediaTypeAAC        HttpMediaType = "audio/aac"
	HttpMediaTypeAVI        HttpMediaType = "video/x-msvideo"
	HttpMediaTypeBinary     HttpMediaType = "application/octet-stream"
	HttpMediaTypeBitmap     HttpMediaType = "image/bmp"
	HttpMediaTypeCSS        HttpMediaType = "text/css"
	HttpMediaTypeCSV        HttpMediaType = "text/csv"
	HttpMediaTypeEPUB       HttpMediaType = "application/epub+zip"
	HttpMediaTypeGZip       HttpMediaType = "application/gzip"
	HttpMediaTypeGIF        HttpMediaType = "image/gif"
	HttpMediaTypeHTML       HttpMediaType = "text/html"
	HttpMediaTypeIcon       HttpMediaType = "image/vnd.microsoft.icon"
	HttpMediaTypeJPEG       HttpMediaType = "image/jpeg"
	HttpMediaTypeJavaScript HttpMediaType = "text/javascript"
	HttpMediaTypeJSON       HttpMediaType = "application/json"
	HttpMediaTypeMP3        HttpMediaType = "audio/mpeg"
	HttpMediaTypeMP4        HttpMediaType = "video/mp4"
	HttpMediaTypeOGGAudio   HttpMediaType = "audio/ogg"
	HttpMediaTypePNG        HttpMediaType = "image/png"
	HttpMediaTypePDF        HttpMediaType = "application/pdf"
	HttpMediaTypePHP        HttpMediaType = "application/php"
	HttpMediaTypeRTF        HttpMediaType = "application/rtf"
	HttpMediaTypeSVG        HttpMediaType = "image/svg+xml"
	HttpMediaTypeSWF        HttpMediaType = "application/x-shockwave-flash"
	HttpMediaTypeTTF        HttpMediaType = "font/ttf"
	HttpMediaTypeText       HttpMediaType = "text/plain"
	HttpMediaTypeWAV        HttpMediaType = "audio/wav"
	HttpMediaTypeWEBMAudio  HttpMediaType = "audio/webm"
	HttpMediaTypeWEBMVideo  HttpMediaType = "video/webm"
	HttpMediaTypeWEBPImage  HttpMediaType = "image/webp"
	HttpMediaTypeWOFF       HttpMediaType = "font/woff"
	HttpMediaTypeWOFF2      HttpMediaType = "font/woff2"
	HttpMediaTypeXHTML      HttpMediaType = "application/xhtml+xml"
	HttpMediaTypeXML        HttpMediaType = "application/xml"
	HttpMediaTypeZip        HttpMediaType = "application/zip"
)

const (
	HttpConnectionClose     HttpConnection = "close"
	HttpConnectionKeepAlive HttpConnection = "keep-alive"
)

const (
	HttpTransferEncodingChunked  HttpTransferEncoding = "chunked"
	HttpTransferEncodingCompress HttpTransferEncoding = "compress"
	HttpTransferEncodingIdentity HttpTransferEncoding = "identity"
	HttpTransferEncodingDeflate  HttpTransferEncoding = "deflate"
	HttpTransferEncodingGZip     HttpTransferEncoding = "gzip"
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
