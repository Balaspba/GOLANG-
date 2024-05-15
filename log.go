package main

import "fmt"

func hotwater() {

}
func singlevariable() {
	fmt.Println("single variable")

	var i int = 7
	var j int
	i = 8
	j = i

	fmt.Println(i)

	i = i * j
	fmt.Println(i)
	i = i + j
	fmt.Println(i)
	i = i % j
	fmt.Println(i)

}
func arraycopy(arr [3]int) {

	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {

			fmt.Println("bala")

		}
	}

}

func arrayvariable() {

	fmt.Println("array variable")

	var arr [3]int = [3]int{10, 20, 30}
	arr[0] = 8
	arr[1] = arr[0]
	fmt.Println(arr)
	arraycopy(arr)
	arr[1] = arr[0] + arr[1]
	arr[0] = arr[1] * arr[2]

	for i := 0; i < 3; i++ {
		arr[1] = 4
		arr[0] = arr[1]
		arr[0] = arr[0] * arr[1]
		arr[i] = arr[i] + arr[i]
	}
	var j int
	for i := 0; i < 3; i++ {
		j = i
		arr[j] = arr[i] + 1

		if j == i {

			arr[j] = arr[i] + 1
		}

	}
}
func main() {
	hotwater()
	singlevariable()
	arrayvariable()

}
