package utils

import (
	"unicode"
)

func WordSplit(c rune) bool {
	return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '-' && c != '\''
}

func SkipTag(tag string) bool {
	return tag == "script" || tag == "span" || tag == "style"
}
