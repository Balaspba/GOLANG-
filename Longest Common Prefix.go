package main

import "fmt"

func LongestCommonPrefix() {
	var strs [3]string = [3]string{"bala", "baju", "bari"}
	if 0 == len(strs) {
		return
	}

	res := ""
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != c {
				return
			}
		}
		res += string(c)
		fmt.Println("Longest Common Prefix", res)

	}
	fmt.Println(LongestCommonPrefix)
	return
}
func main() {
	LongestCommonPrefix()

}
