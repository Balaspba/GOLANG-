package main

import (
	"fmt"
)

func main() {

	number := 10

	if number%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	language := "Golang"

	switch language {
	case "Python":
		fmt.Println("You selected Python")
	case "Java":
		fmt.Println("You selected Java")
	case "Golang":
		fmt.Println("You selected Golang")
	default:
		fmt.Println("Language not found")
	}

	fmt.Println("Printing numbers from 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("Nested loop example:")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d,%d ", i, j)
		}
		fmt.Println()
	}
}
