package main

import "fmt"

func Equilibriumindex() {
	arr := []int{1, 2, 3, 4, 5, 6}
	n := len(arr)

	for i := 0; i < n; i++ {
		leftsum := 0
		rightsum := 0

		for j := 0; j < i; j++ {
			leftsum += arr[j]
		}

		for j := i + 1; j < n; j++ {
			rightsum += arr[j]
		}

		if leftsum == rightsum {
			fmt.Println("Equilibrium index found at:", i)
			return
		}
	}

	fmt.Println("No equilibrium index found")
}

func main() {
	Equilibriumindex()
}
