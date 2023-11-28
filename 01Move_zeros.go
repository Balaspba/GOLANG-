package main

import (
	"fmt"
)

func movezero() {
	var nums [5]int = [5]int{0, 22, 77, 99, 88}
	var k int
	for i := range nums {
		if nums[i] != 0 {
			if k != i {

				fmt.Println("moves zeros", nums)
			}
			k++
		}

	}

}
func main() {
	movezero()

}
