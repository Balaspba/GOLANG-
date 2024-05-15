package main

import "fmt"

func maxProfit() {
	var prices [5]int = [5]int{1, 2, 3, 4, 5}
	n := len(prices)

	if 0 == n || 1 == n {
		return
	}

	var (
		res      int
		minPrice = prices[0]
	)
	for i := 1; i < n; i++ {
		if prices[i] < prices[i-1] {
			fmt.Println("offer")
		}
		if i == n-1 {
			fmt.Println("no offer")
		}
	}
	fmt.Println(res)
	return
}
func main() {
	maxProfit()
}
