package http

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Method string
type StatusCode int
type Version string
type MediaType string
type Header string
type ConnectionHeader string
type TransferEncodingHeader string
type ExpectHeader string

const (
	MethodGet     Method = "GET"
	MethodHead    Method = "HEAD"
	MethodPost    Method = "POST"
	MethodPut     Method = "PUT"
	MethodDelete  Method = "DELETE"
	MethodConnect Method = "CONNECT"
	MethodOptions Method = "OPTIONS"
	MethodTrace   Method = "TRACE"
)

const (
	StatusContinue StatusCode = iota + 100
	StatusSwitchingProtocols
	StatusProcessing
)

const (
	StatusOK StatusCode = iota + 200
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
	StatusMultipleChoices StatusCode = iota + 300
	StatusMovedPermanently
	StatusFound
	StatusSeeOther
	StatusNotModified
	StatusUseProxy
	StatusTemporaryRedirect
	StatusPermanentRedirect
)

const (
	StatusBadRequest StatusCode = iota + 400
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
	StatusInternalServerError StatusCode = iota + 500
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
	Version09 Version = "HTTP/0.9"
	Version10 Version = "HTTP/1.0"
	Version11 Version = "HTTP/1.1"
)

const (
	MediaTypeAAC        MediaType = "audio/aac"
	MediaTypeAVI        MediaType = "video/x-msvideo"
	MediaTypeBinary     MediaType = "application/octet-stream"
	MediaTypeBitmap     MediaType = "image/bmp"
	MediaTypeCSS        MediaType = "text/css"
	MediaTypeCSV        MediaType = "text/csv"
	MediaTypeEPUB       MediaType = "application/epub+zip"
	MediaTypeGZip       MediaType = "application/gzip"
	MediaTypeGIF        MediaType = "image/gif"
	MediaTypeHTML       MediaType = "text/html"
	MediaTypeHTTP       MediaType = "message/http"
	MediaTypeIcon       MediaType = "image/vnd.microsoft.icon"
	MediaTypeJPEG       MediaType = "image/jpeg"
	MediaTypeJavaScript MediaType = "text/javascript"
	MediaTypeJSON       MediaType = "application/json"
	MediaTypeMP3        MediaType = "audio/mpeg"
	MediaTypeMP4        MediaType = "video/mp4"
	MediaTypeOGGAudio   MediaType = "audio/ogg"
	MediaTypePNG        MediaType = "image/png"
	MediaTypePDF        MediaType = "application/pdf"
	MediaTypePHP        MediaType = "application/php"
	MediaTypeRTF        MediaType = "application/rtf"
	MediaTypeSVG        MediaType = "image/svg+xml"
	MediaTypeSWF        MediaType = "application/x-shockwave-flash"
	MediaTypeTTF        MediaType = "font/ttf"
	MediaTypeText       MediaType = "text/plain"
	MediaTypeWAV        MediaType = "audio/wav"
	MediaTypeWEBMAudio  MediaType = "audio/webm"
	MediaTypeWEBMVideo  MediaType = "video/webm"
	MediaTypeWEBPImage  MediaType = "image/webp"
	MediaTypeWOFF       MediaType = "font/woff"
	MediaTypeWOFF2      MediaType = "font/woff2"
	MediaTypeXHTML      MediaType = "application/xhtml+xml"
	MediaTypeXML        MediaType = "application/xml"
	MediaTypeZip        MediaType = "application/zip"
)

const (
	HeaderHost             Header = "host"
	HeaderConnection       Header = "connection"
	HeaderContentLength    Header = "content-length"
	HeaderContentType      Header = "content-type"
	HeaderTransferEncoding Header = "transfer-encoding"
	HeaderTE               Header = "te"
	HeaderUpgrade          Header = "upgrade"
	HeaderVia              Header = "via"
	HeaderLastModified     Header = "last-modified"
	HeaderETag             Header = "etag"
	HeaderExpect           Header = "expect"
	HeaderServer           Header = "server"
	HeaderDate             Header = "date"
	HeaderAllow            Header = "allow"
)

const (
	ConnectionHeaderClose     ConnectionHeader = "close"
	ConnectionHeaderKeepAlive ConnectionHeader = "keep-alive"
)

const (
	TransferEncodingHeaderChunked  TransferEncodingHeader = "chunked"
	TransferEncodingHeaderCompress TransferEncodingHeader = "compress"
	TransferEncodingHeaderIdentity TransferEncodingHeader = "identity"
	TransferEncodingHeaderDeflate  TransferEncodingHeader = "deflate"
	TransferEncodingHeaderGZip     TransferEncodingHeader = "gzip"
)

const ExpectHeaderContinue ExpectHeader = "100-continue"

type Form int
type Scheme string

const (
	FormOrigin Form = iota
	FormAbsolute
	FormAuthority
	FormAsterisk
)

const (
	SchemeHttp  Scheme = "http"
	SchemeHttps Scheme = "https"
)

func encodePercent(str string) string {
	encoded := ""
	for _, char := range str {
		if (char < 0x21 || char > 0x7E) && char >= 0 && char <= math.MaxUint8 {
			encoded += "%" + fmt.Sprintf("%02x", char)
		} else {
			encoded += string(char)
		}
	}
	return encoded
}

func decodePercent(str string) string {
	if len(str) < 3 {
		return str
	}

	decoded := ""
	for index, char := 0, str[0]; index < len(str); index++ {
		if char == '%' && index < len(str)-2 {
			if char, err := strconv.ParseInt(str[index+1:index+3], 16, 16); err == nil {
				decoded += string(char)
			}
			index += 2
		} else {
			decoded += string(char)
		}
		if index < len(str)-1 {
			char = str[index+1]
		}
	}
	return decoded
}

func formatTimeGMT(t time.Time) string {
	return t.UTC().Format(time.RFC1123[:len(time.RFC1123)-3]) + "GMT"
}

func normalizeCase(str string) string {
	return strings.ToLower(str)
}

func isValidHeaderValue(str string) bool {
	for _, char := range str {
		if (char < 0x21 || char > 0x7E) && char != ' ' && char != '\t' {
			return false
		}
	}
	return true
}

func isVisibleString(str string) bool {
	for _, char := range str {
		if char < 0x21 || char > 0x7E {
			return false
		}
	}
	return true
}
