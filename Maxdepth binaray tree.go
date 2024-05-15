package main

import (
	"fmt"
)

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func maxDepth() {
	var root *TreeNode1 = (&TreeNode1{Val: 1, Left: &TreeNode1{Val: 9, Right: &TreeNode1{Val: 20}}})

	if root == nil {
		return
	}

	left := root.Left
	right := root.Right

	var max int
	if left > right {
		max = left
	} else {
		max = right
	}
	fmt.Println("Given a binary tree, find its maximum depth.", max+1)
	return

}
func main() {
	maxDepth()

}
