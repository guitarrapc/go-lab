package main

import (
	"testing"
	"unicode/utf8"
)

var s = "hogemogehogemogehogemoge世界世界世界a"

func BenchmarkGetLastRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getLastRune(s, 3)
	}
}

func BenchmarkGetLastRune2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getLastRune2(s, 3)
	}
}

func getLastRune(s string, c int) string {
	j := len(s)
	for i := 0; i < c && j > 0; i++ {
		_, size := utf8.DecodeLastRuneInString(s[:j])
		j -= size
	}
	return s[j:]
}

func getLastRune2(s string, c int) string {
	r := []rune(s)
	if c > len(r) {
		c = len(r)
	}
	return string(r[len(r)-c:])
}
