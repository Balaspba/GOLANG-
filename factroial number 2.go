package main

import "fmt"

func fact(n int) {

	x := 1
	for i := 1; i <= n; i++ {
		x *= i
		fmt.Println(x)
	}
}
func main() {
	fact(6)
}
