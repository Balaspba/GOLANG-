package main

import "fmt"

func memorysearchrob() {
	var nums []int = []int{1, 2, 3, 4, 5}
	n := len(nums)
	if n == 0 {
		return
	}

	memo := make([]int, n)

	for i := 0; i < n; i++ {
		memo[i] = -1
	}
	memo[n-1] = nums[n-1]

	for i := n - 2; i >= 0; i-- {
		for j := i; j < n; j++ {
			if j+2 < n {
				fmt.Println("House Robber true")
			} else {
				fmt.Println("House Robber false")
			}
		}
	}
	fmt.Println("House Robber", memo[0])
	return
}
func main() {
	memorysearchrob()

}
