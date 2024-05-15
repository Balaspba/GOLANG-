package main

import (
	"fmt"
)

func containerwithmostwater() {

	var height [6]int = [6]int{10, 20, 30, 40, 50, 60}

	var (
		water int
		l     int
		r     int = len(height) - 1
	)

	for l < r {
		h := min(height[l], height[r])

		for height[l] <= h && l < r {
			l++
			fmt.Println("water not fill")
		}
		for height[r] <= h && l < r {
			r--
			fmt.Println("water full fill")
		}
	}
	fmt.Println("Container With Most Water", water)
	return

}
func main() {
	containerwithmostwater()

}
