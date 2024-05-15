package main

import (
	"fmt"
	"math"
)

func findSecondLargest(arr []int) int {
	max := math.MinInt64
	secondMax := math.MinInt64

	for _, num := range arr {
		if num > max {
			secondMax = max
			max = num
		} else if num > secondMax && num < max {
			secondMax = num
		}
	}

	return secondMax
}

func main() {
	arr := []int{5, 2, 9, 1, 7, 3}
	secondLargest := findSecondLargest(arr)
	fmt.Println("The second largest element is:", secondLargest)
}
