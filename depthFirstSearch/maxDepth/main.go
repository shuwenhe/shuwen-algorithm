package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("depth first search")
	r7 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	l15 := &TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}

	l9 := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}

	r20 := &TreeNode{
		Val:   9,
		Left:  l15,
		Right: r7,
	}

	root := &TreeNode{
		Val:   3,
		Left:  l9,
		Right: r20,
	}

	fmt.Println(maxDepth(root))
}
