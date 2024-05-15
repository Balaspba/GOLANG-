package main

import "fmt"

func fill() [][]int {
	base := [][]int{{1, 2, 2}, {1, 2, 1}, {1, 1, 1}}

	numRows := len(base)
	for row := 1; row < numRows; row++ {
		for col := 1; col < len(base[row])-1; col++ {
			base[row][col] = base[row-1][col-1] + base[row-1][col]
		}
	}
	fmt.Println(base)
	return base
}
func main() {

	fill()
}
