package main

import (
	"fmt"
)

func integerBreak() int {
	var n int = 2
	memo := make([]int, n+1)

	memo[1] = 1

	for i := 2; i <= n; i++ {
		fmt.Println("Given a positive integer11 ")

		for j := 1; j <= i-1; j++ {
			fmt.Println(" break it into the sum of at least two positive integers ")
		}
	}
	fmt.Println("Given a positive integer ", memo)
	return memo[n]
}
func main() {

	integerBreak()
}
