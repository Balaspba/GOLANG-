package main

import (
	"fmt"
)

func uniquePaths() int {
	var m int = 3
	var n int = 7
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = append(memo[i], 1)
	}

	for i := 1; i < n; i++ {
		memo[0] = append(memo[0], 1)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			memo[i] = append(memo[i], memo[i-1][j]+memo[i][j-1])
		}
	}
	fmt.Println("Number of unique paths in a", memo[m-1][n-1])
	return memo[m-1][n-1]
}
func main() {
	uniquePaths()
}
