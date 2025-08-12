// Package string is a library to extend the string functionality of the Go programming language
// It orients itself on the php base library and provides similar functions
package string

import array "girado73/goLib/baselib/array"

func Chars(s string) []rune {
	var chars []rune
	for _, c := range s {
		chars = append(chars, c)
	}
	return chars
}

func ChunkSplit(s string, length int, end string) string {
	if length <= 0 {
		return s
	}

	var result string
	for i := 0; i < len(s); i += length {
		endIndex := min(i+length, len(s))
		result += s[i:endIndex] + end
	}
	return result[:len(result)-len(end)] // Remove the last added end
}

func JoinRunes(s []rune) string {
	return string(s)
}

func Trim(s string) string {
	chars := Chars(s)

	return string(array.Filter(chars, func(c rune) bool { return c != ' ' }))
}

func StrToUpper(s string) string {
	chars := Chars(s)
	for i, c := range chars {
		if c >= 'a' && c <= 'z' {
			chars[i] = c - ('a' - 'A')
		}
	}
	return JoinRunes(chars)
}

func StrToLower(s string) string {
	chars := Chars(s)
	for i, c := range chars {
		if c >= 'A' && c <= 'Z' {
			chars[i] = c + ('a' - 'A')
		}
	}
	return JoinRunes(chars)
}
