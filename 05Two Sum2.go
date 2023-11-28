package main

import (
	"fmt"
)

func twoSum2() {
	var numbers [8]int = [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	target := 15
	n := len(numbers)
	var (
		l = 0
		r = n - 1
	)
	fmt.Println("twosum2", numbers)
	for l < r {

		if numbers[l]+numbers[r] == target {

			fmt.Println("Two sum2", l+1, r+1)
			return

		} else if numbers[l]+numbers[r] < target {
			l++
			fmt.Println("bala")
		} else {
			r--

		}

	}

}

func main() {

	twoSum2()
}
