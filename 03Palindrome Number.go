package main

import (
	"fmt"
)

func Palindrome() {
	var originalNumber int = 12321
	var reverseNumber int = 0
	var temp int = originalNumber

	for temp > 0 {
		var remainder = temp % 10
		reverseNumber = (reverseNumber * 10) + remainder
		temp = temp / 10
	}

	if originalNumber == reverseNumber {
		fmt.Println("The given number is Palindrome")
	} else {
		fmt.Println("The given number is NOT a Palindrome")
	}
}

func main() {

	Palindrome()
}
