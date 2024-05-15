package main

import (
	"fmt"
	"strconv"
)

func decimalToBinary(decimal int) string {
	var binary []int

	for decimal > 0 {
		remainder := decimal % 2
		binary = append(binary, remainder)
		decimal = decimal / 2
	}

	for i, j := 0, len(binary)-1; i < j; i, j = i+1, j-1 {
		binary[i], binary[j] = binary[j], binary[i]
	}

	binaryString := ""
	for _, digit := range binary {
		binaryString += strconv.Itoa(digit)
	}

	return binaryString
}

func main() {
	decimal := 35
	binary := decimalToBinary(decimal)
	fmt.Printf("Binary representation of %d is %s\n", decimal, binary)
}
