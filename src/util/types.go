package util

type HttpMethod string
type HttpStatusCode int
type HttpVersion float64
type HttpMediaType string

const (
	HttpMethodGet     HttpMethod = "GET"
	HttpMethodHead               = "HEAD"
	HttpMethodPost               = "POST"
	HttpMethodPut                = "PUT"
	HttpMethodDelete             = "DELETE"
	HttpMethodConnect            = "CONNECT"
	HttpMethodOptions            = "OPTIONS"
	HttpMethodTrace              = "TRACE"
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
	HttpVersion09 HttpVersion = 0.9
	HttpVersion10             = 1.0
	HttpVersion11             = 1.1
	HttpVersion20             = 2.0
)

const (
	HttpMediaTypeAAC        HttpMediaType = "audio/aac"
	HttpMediaTypeAVI                      = "video/x-msvideo"
	HttpMediaTypeBinary                   = "application/octet-stream"
	HttpMediaTypeBitmap                   = "image/bmp"
	HttpMediaTypeCSS                      = "text/css"
	HttpMediaTypeCSV                      = "text/csv"
	HttpMediaTypeEPUB                     = "application/epub+zip"
	HttpMediaTypeGZip                     = "application/gzip"
	HttpMediaTypeGIF                      = "image/gif"
	HttpMediaTypeHTML                     = "text/html"
	HttpMediaTypeIcon                     = "image/vnd.microsoft.icon"
	HttpMediaTypeJPEG                     = "image/jpeg"
	HttpMediaTypeJavaScript               = "text/javascript"
	HttpMediaTypeJSON                     = "application/json"
	HttpMediaTypeMP3                      = "audio/mpeg"
	HttpMediaTypeMP4                      = "video/mp4"
	HttpMediaTypeOGGAudio                 = "audio/ogg"
	HttpMediaTypePNG                      = "image/png"
	HttpMediaTypePDF                      = "application/pdf"
	HttpMediaTypePHP                      = "application/php"
	HttpMediaTypeRTF                      = "application/rtf"
	HttpMediaTypeSVG                      = "image/svg+xml"
	HttpMediaTypeSWF                      = "application/x-shockwave-flash"
	HttpMediaTypeTTF                      = "font/ttf"
	HttpMediaTypeText                     = "text/plain"
	HttpMediaTypeWAV                      = "audio/wav"
	HttpMediaTypeWEBMAudio                = "audio/webm"
	HttpMediaTypeWEBMVideo                = "video/webm"
	HttpMediaTypeWEBPImage                = "image/webp"
	HttpMediaTypeWOFF                     = "font/woff"
	HttpMediaTypeWOFF2                    = "font/woff2"
	HttpMediaTypeXHTML                    = "application/xhtml+xml"
	HttpMediaTypeXML                      = "application/xml"
	HttpMediaTypeZip                      = "application/zip"
)

type UriForm int
type UriScheme string

const (
	UriFormOrigin UriForm = iota
	UriFormAbsolute
	UriFormAuthority
	UriFormAsterisk
)

const (
	UriSchemeHttp  = "http"
	UriSchemeHttps = "https"
)
