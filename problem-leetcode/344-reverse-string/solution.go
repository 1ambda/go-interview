package leetcode

import (
	"unicode"
	"unicode/utf8"
)

// http://rosettacode.org/wiki/Reverse_a_string#Go
// support unicode combining characters version too.
func reverseString(s string) string {
	if s == "" {
		return ""
	}
	p := []rune(s)
	r := make([]rune, len(p))
	start := len(r)
	for i := 0; i < len(p); {
		// quietly skip invalid UTF-8
		if p[i] == utf8.RuneError {
			i++
			continue
		}
		j := i + 1
		for j < len(p) && (unicode.Is(unicode.Mn, p[j]) ||
			unicode.Is(unicode.Me, p[j]) || unicode.Is(unicode.Mc, p[j])) {
			j++
		}
		for k := j - 1; k >= i; k-- {
			start--
			r[start] = p[k]
		}
		i = j
	}
	return string(r[start:])
}

// doesn't support unicode
func reverseStringSimple(s string) string {
	runes :=[]rune(s)

	for i, j := 0, len(runes) - 1; i <j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
