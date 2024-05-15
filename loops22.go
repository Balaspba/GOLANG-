package main

import "fmt"

func main() {

	fmt.Println("loops start")
	arr := [4]int{0, 1, 2, 3}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
		for j := 0; j < i-1; j++ {
			if j == 0 {
				fmt.Println("iner loop")
			} else if j < 3 {
				fmt.Println("outter loop")
			}
			fmt.Println("loops end")
		}

		switch i {
		case 0, 1, 2:
			fmt.Println("value")

		default:
			fmt.Println("nothing")
		}

	}
}
