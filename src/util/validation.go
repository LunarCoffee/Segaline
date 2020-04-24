package util

func IsValidHeaderValue(str string) bool {
	for _, char := range str {
		if !isVisibleChar(char) && char != ' ' && char != '\t' {
			return false
		}
	}
	return true
}

func IsVisibleString(str string) bool {
	for _, char := range str {
		if !isVisibleChar(char) {
			return false
		}
	}
	return true
}

func isVisibleChar(char rune) bool {
	return char >= 0x21 && char <= 0x7E
}
