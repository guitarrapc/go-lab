package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// which is more better for memory allocation?
	s := "hogemogehogemogehogemoge世界世界世界a"
	getLastRune(s, 3)
	getLastRune2(s, 3)
}

func getLastRune(s string, c int) {
	// DecodeLastRuneInString
	j := len(s)
	for i := 0; i < c && j > 0; i++ {
		_, size := utf8.DecodeLastRuneInString(s[:j])
		j -= size
	}
	lastByRune := s[j:]
	fmt.Println(lastByRune)
}

func getLastRune2(s string, c int) {
	// string -> []rune
	r := []rune(s)
	lastByRune := string(r[len(r)-c:])
	fmt.Println(lastByRune)
}
