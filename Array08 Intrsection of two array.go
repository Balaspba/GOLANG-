package main

import "fmt"

func intersection() []int {

	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	set := make(map[int]bool)
	result := []int{}

	for _, num := range nums1 {
		set[num] = true
	}

	for _, num := range nums2 {
		if set[num] {
			result = append(result, num)

			delete(set, num)
		}
	}
	fmt.Println("Intersection of the two arrays:", result)
	return result
}

func main() {

	intersection()

}
