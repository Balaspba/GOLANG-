package main

import (
	"fmt"
)

func reverseinteger() {

	var x int = 123456
	if 0 == x {

		return

	}

	isPositive := true
	if x < 0 {
		isPositive = false
		x *= -1
		fmt.Println(x)
	}

	res := 0

	for x > 0 {
		res = res*10 + x%10
		x /= 10
		fmt.Println(x)
	}

	if !isPositive {
		res *= -1
		fmt.Println("isPositive")
	}

	if res < -1<<31 || res > (1<<31)-1 {
		fmt.Println(0)
		return

	}

	fmt.Println(res)

}
func main() {

	reverseinteger()

}
