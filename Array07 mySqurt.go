package main

import "fmt"

func mysqurt() int {
	var x int = 8
	var out int = 0
	for out*out < x {
		out++
	}
	if out*out > x {
		out--
	}
	fmt.Println("My squrt number is", out)
	return 0
}
func main() {
	mysqurt()
}
