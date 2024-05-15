package main

import (
	"fmt"
)

func intersect() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	record := make(map[int]int)
	res := make([]int, 0)
	for _, num := range nums1 {
		if _, ok := record[num]; !ok {
			record[num] = 1
		} else {
			record[num]++
		}
	}

	for _, num := range nums2 {
		if count, ok := record[num]; ok && count > 0 {
			res = append(res, num)
			record[num]--
		}
	}
	fmt.Println("Intersection of Two Arrays II", res)
	return
}
func main() {
	intersect()
}
