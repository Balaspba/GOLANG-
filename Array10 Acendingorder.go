package main

import "fmt"

func acedce(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{2, 1, 4, 6, 3, 5}
	res := acedce(arr)
	fmt.Println(res)
}
