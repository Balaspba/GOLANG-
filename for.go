package main

import (
	"fmt"
)

func main() {

	nums := [3]string{"one", "two", "three"}
	for idx, num := range nums {
		fmt.Printf("%d: %s\n", idx, num)
	}
	for idx := range nums {
		fmt.Println(idx)
	}
	for range nums {
		fmt.Println("tick")
	}
}
