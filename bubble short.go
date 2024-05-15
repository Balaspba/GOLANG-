package main

import "fmt"

func bubbleSort(numbers []int) []int {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}

func main() {
	numbers := []int{64, 25, 12, 22, 11}
	sorted := bubbleSort(numbers)
	fmt.Println("Sorted array:", sorted)
}
