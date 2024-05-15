package main

import "fmt"

func main() {

	fmt.Println("Loops")

	for i := 0; i < 5; i++ {

		fmt.Println(" outerloop starts ")

		for j := 0; j < i; j++ {

			fmt.Println(" innerloop  starts ")

			if j == 4 {
				fmt.Println(" innerloop  staisfies ")
			}

		}

		fmt.Println(" outerloop ends ")

	}
}
