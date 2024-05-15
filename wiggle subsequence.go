package main

import (
	"fmt"
)

func wiggleMaxLength() {

	var nums [2]int = [2]int{1, 2}
	n := len(nums)

	if n == 1 || n == 0 {
		return
	}

	up := make([]int, n)
	for i := 0; i < n; i++ {
		up[i] = 1
	}
	down := make([]int, n)
	copy(down, up)

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {

			} else if nums[i] < nums[j] {

			}
		}
	}
	fmt.Println("The entire sequence is a wiggle sequence.", up[n-1], down[n-1])
	return
}
func main() {

	wiggleMaxLength()
}
