package main

import "fmt"

func main() {

	size := 10

	oddNumbers := make([]int, size)

	for i := 0; i < size; i++ {
		oddNumbers[i] = 2*i + 2
	}

	fmt.Println(oddNumbers)

}
