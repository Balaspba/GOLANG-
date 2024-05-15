package main

import (
	"fmt"
	"sort"
)

func missingPositive() int {
	arr := []int{1, 2, 5, 3, -4, 6, 8, 9}
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			if i == 0 || arr[i] != arr[i-1]+1 {
				return arr[i] + 1
			}
		}
	}
	return 0
}

func main() {
	result := missingPositive()
	fmt.Println("Missing positive integer:", result)
}
