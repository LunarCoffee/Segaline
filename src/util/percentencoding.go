package util

import (
	"fmt"
	"math"
	"strconv"
)

func PercentEncode(str string) string {
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

func PercentDecode(str string) string {
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
