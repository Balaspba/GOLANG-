package main

import "fmt"

func rotateArray() []int {

	arr := []int{1, 2, 3, 4, 5}
	k := 1

	n := len(arr)

	k = k % n

	rotated := make([]int, n)

	for i := 0; i < n; i++ {
		rotated[(i+k)%n] = arr[i]

	}
	fmt.Printf("Rotated Array by  %d positions:", k, rotated)

	return rotated
}

func main() {

	rotateArray()

}
