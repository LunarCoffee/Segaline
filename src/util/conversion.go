package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

func ContentTypeByExt(ext string) HttpMediaType {
	switch ext {
	case "aac":
		return HttpMediaTypeAAC
	case "avi":
		return HttpMediaTypeAVI
	case "bmp":
		return HttpMediaTypeBitmap
	case "css":
		return HttpMediaTypeCSS
	case "csv":
		return HttpMediaTypeCSV
	case "epub":
		return HttpMediaTypeEPUB
	case "gz":
		return HttpMediaTypeGZip
	case "gif":
		return HttpMediaTypeGIF
	case "htm", "html":
		return HttpMediaTypeHTML
	case "ico":
		return HttpMediaTypeIcon
	case "jpg", "jpeg":
		return HttpMediaTypeJPEG
	case "js":
		return HttpMediaTypeJavaScript
	case "json":
		return HttpMediaTypeJSON
	case "mp3":
		return HttpMediaTypeMP3
	case "mp4":
		return HttpMediaTypeMP4
	case "oga":
		return HttpMediaTypeOGGAudio
	case "png":
		return HttpMediaTypePNG
	case "pdf":
		return HttpMediaTypePDF
	case "php":
		return HttpMediaTypePHP
	case "rtf":
		return HttpMediaTypeRTF
	case "svg":
		return HttpMediaTypeSVG
	case "swf":
		return HttpMediaTypeSWF
	case "ttf":
		return HttpMediaTypeTTF
	case "txt":
		return HttpMediaTypeText
	case "wav":
		return HttpMediaTypeWAV
	case "weba":
		return HttpMediaTypeWEBMAudio
	case "webm":
		return HttpMediaTypeWEBMVideo
	case "webp":
		return HttpMediaTypeWEBPImage
	case "woff":
		return HttpMediaTypeWOFF
	case "woff2":
		return HttpMediaTypeWOFF2
	case "xhtml":
		return HttpMediaTypeXHTML
	case "xml":
		return HttpMediaTypeXML
	case "zip":
		return HttpMediaTypeZip
	}
	return HttpMediaTypeBinary
}
