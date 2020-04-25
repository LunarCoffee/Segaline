package util

import (
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func EncodePercent(str string) string {
	encoded := ""
	for _, char := range str {
		if !isVisibleChar(char) && char >= 0 && char <= math.MaxUint8 {
			encoded += "%" + fmt.Sprintf("%02x", char)
		} else {
			encoded += string(char)
		}
	}
	return encoded
}

func DecodePercent(str string) string {
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

func NormalizeCase(str string) string {
	return strings.ToLower(str)
}

func GetETag(content []byte) string {
	sha := sha1.New()
	sha.Write(content)
	return base32.HexEncoding.EncodeToString(sha.Sum(nil))
}

func FormatTimeGMT(time time.Time) string {
	time = time.UTC()
	return time.Format(ResponseTimeFormat) + " GMT"
}

func FormatErrorTemplate(template string, status HttpStatusCode) string {
	statusReplaced := strings.ReplaceAll(template, "{statusCode}", strconv.Itoa(int(status)))
	return strings.ReplaceAll(statusReplaced, "{serverInfo}", ServerNameVersion)
}

func ContentTypeByExt(ext string) HttpMediaType {
	switch ext {
	case "aac":
		return MediaTypeAAC
	case "avi":
		return MediaTypeAVI
	case "bmp":
		return MediaTypeBitmap
	case "css":
		return MediaTypeCSS
	case "csv":
		return MediaTypeCSV
	case "epub":
		return MediaTypeEPUB
	case "gz":
		return MediaTypeGZip
	case "gif":
		return MediaTypeGIF
	case "htm", "html":
		return MediaTypeHTML
	case "ico":
		return MediaTypeIcon
	case "jpg", "jpeg":
		return MediaTypeJPEG
	case "js":
		return MediaTypeJavaScript
	case "json":
		return MediaTypeJSON
	case "mp3":
		return MediaTypeMP3
	case "mp4":
		return MediaTypeMP4
	case "oga":
		return MediaTypeOGGAudio
	case "png":
		return MediaTypePNG
	case "pdf":
		return MediaTypePDF
	case "php":
		return MediaTypePHP
	case "rtf":
		return MediaTypeRTF
	case "svg":
		return MediaTypeSVG
	case "swf":
		return MediaTypeSWF
	case "ttf":
		return MediaTypeTTF
	case "txt":
		return MediaTypeText
	case "wav":
		return MediaTypeWAV
	case "weba":
		return MediaTypeWEBMAudio
	case "webm":
		return MediaTypeWEBMVideo
	case "webp":
		return MediaTypeWEBPImage
	case "woff":
		return MediaTypeWOFF
	case "woff2":
		return MediaTypeWOFF2
	case "xhtml":
		return MediaTypeXHTML
	case "xml":
		return MediaTypeXML
	case "zip":
		return MediaTypeZip
	}
	return MediaTypeBinary
}
