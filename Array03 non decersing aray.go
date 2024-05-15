package main

import "fmt"

func nondecersingarray() {

	var n int = 5
	arr := make([]int, n)
	arr[0] = 1
	for i := 1; i < n; i++ {

		arr[i] = arr[i-1] + i

	}
	fmt.Println(arr)
	return

}
func main() {
	nondecersingarray()
}
