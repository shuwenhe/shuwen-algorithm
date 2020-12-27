package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func main() {
	l1 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	r3 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	root := &TreeNode{
		Val:   2,
		Left:  l1,
		Right: r3,
	}
	fmt.Println(invertTree(root))
	fmt.Println(invertTree(root).Left.Val)
	fmt.Println(invertTree(root).Right.Val)
}
