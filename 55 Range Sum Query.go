package main

import "fmt"

type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	nA := NumArray{sum: make([]int, 1)}
	for _, num := range nums {

		fmt.Println("Constructor make NumArray instance.", nA, num)
	}
	return nA
}

func main() {
	var nums []int = []int{-2, 0, 3, -5, 2, -1}
	Constructor(nums)
}
