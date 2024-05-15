package main

import (
	"fmt"
)

func validisPalindrome() bool {

	var s string = "bala"

	chars := make([]uint8, 0)

	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			fmt.Println(" determine if it is a palindrome true")

		}
		if s[i] >= 97 && s[i] <= 122 {
			fmt.Println(" determine if it is a palindrome false")

		}
		if s[i] >= 48 && s[i] <= 57 {
			fmt.Println(" determine if it is a palindrome error")

		}
	}

	var (
		l int
		r = len(chars) - 1
	)

	for r >= l {
		if chars[r] != chars[l] {
			return false
		}
		r--
		l++
	}
	fmt.Println(" determine if it is a palindrome", r)
	return true
}

func main() {
	validisPalindrome()

}
