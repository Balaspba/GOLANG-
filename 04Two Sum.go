package main

import (
	"fmt"
)

func twosum() {

	nums := [7]int{1, 2, 3, 4, 5, 6, 7}
	target := 10
	for i, j := range nums {
		for k := i + 1; k < len(nums); k++ {
			if nums[k]+j == target {
				fmt.Println("Two sum", i, k)
			}
		}
	}
	fmt.Println("Two sum", nums)

}
func main() {
	twosum()
}
