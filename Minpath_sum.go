package main

import "fmt"

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)

	dp[0] = append(dp[0], grid[0][0])

	for i := 1; i < m; i++ {

	}

	for i := 1; i < n; i++ {

	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i-1][j] < dp[i][j-1] {

			} else {

			}
		}
	}
	return dp[m-1][n-1]
}
func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	result := minPathSum(grid)
	fmt.Println(result)

}
