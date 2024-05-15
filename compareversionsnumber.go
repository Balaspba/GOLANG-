package main

import (
	"fmt"
	"strings"
)

func compareVersion() {

	var version1 string = "bala"
	var version2 string = "hari"
	var (
		v1 = strings.Split(version1, ".")
		v2 = strings.Split(version2, ".")
	)
	for i := 0; i < len(v1) || i < len(v2); i++ {
		var v1N, v2N int
		if i < len(v1) {
			fmt.Println(" Compare Version Numbers 1")

		}
		if i < len(v2) {
			fmt.Println(" Compare Version Numbers2")
		}

		if v1N > v2N {
			fmt.Println(" Compare Version Numbers3")
			return
		} else if v1N < v2N {
			return
		}
	}
	fmt.Println(" Compare Version Numbers4", v1, v2)
	return
}
func main() {
	compareVersion()

}
