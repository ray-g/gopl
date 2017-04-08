package unispace

import (
	"unicode"
	"unicode/utf8"
)

func uniSpace(s []byte) []byte {
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRune(s[i:])
		if unicode.IsSpace(r) {
			copy(s[i+1:], s[i+size:])
			s[i] = ' '
			s = s[:len(s)-size+1] // len(slice)-len(unicode space)+len(ASCII space)
			i++
		} else {
			i += size
		}
	}
	return s
}
