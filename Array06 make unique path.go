package main

import "fmt"

func main() {

	input := []int{1, 2, 3, 4, 2, 3, 5, 6, 1}

	uniqueMap := make(map[int]bool)

	for _, num := range input {

		uniqueMap[num] = true
	}

	uniqueArray := make([]int, 0, len(uniqueMap))

	for num := range uniqueMap {
		uniqueArray = append(uniqueArray, num)
	}

	fmt.Println(uniqueArray)
}
