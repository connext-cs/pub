package utils

import "strings"

func CheckStrEmpty(str ...string) bool {
	for _, s := range str {
		if s == "" {
			return true
		}
	}
	return false
}

func CheckIntEmpty(ints ...int) bool {
	for _, v := range ints {
		if v == 0 {
			return true
		}
	}
	return false
}

func StrimSpaceAndEof(str_byte []byte) (result string) {
	result = strings.Replace(string(str_byte), " ", "", -1)
	result = strings.Replace(result, "\n", "", -1)
	return
}
