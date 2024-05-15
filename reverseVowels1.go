package main

import (
	"fmt"
)

func isVowel(char byte) bool {

	vowels := "aeiouAEIOU"
	for _, v := range vowels {
		if char == byte(v) {
			return true
		}
	}
	return false
}
func reverseVowels() {
	s := "hello"

	var (
		l int
		r = len(s) - 1
	)

	for r > l {

		for r >= 0 && !isVowel(s[r]) {
			r--
		}

		for l < len(s) && !isVowel(s[l]) {
			l++
		}

		if l >= r {
			break
		}
		charL := s[l]
		charR := s[r]

		s = s[:r] + string(charL) + s[r+1:]
		s = s[:l] + string(charR) + s[l+1:]

		r--
		l++
	}
	fmt.Println("reverseVowels", s)
	return
}
func main() {
	reverseVowels()
}
