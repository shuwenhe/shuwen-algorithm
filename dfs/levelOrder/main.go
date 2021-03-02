// 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
// 例如:
// 给定二叉树: [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回其层次遍历结果：
// [
//   [3],
//   [9,20],
//   [15,7]
// ]

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t15 := &TreeNode{
		Val: 15,
	}
	t7 := &TreeNode{
		Val: 7,
	}
	t20 := &TreeNode{
		Val:   20,
		Left:  t15,
		Right: t7,
	}
	t9 := &TreeNode{
		Val: 9,
	}
	root := &TreeNode{
		Val:   3,
		Left:  t9,
		Right: t20,
	}
	arr := levelOrder(root)
	for k, v := range arr {
		fmt.Printf("level %d = %d\n", k, v)
	}
}

var result [][]int

func levelOrder(root *TreeNode) [][]int {
	result = make([][]int, 0)
	if root == nil {
		return result
	}
	dfs(root, 0)
	return result
}

func dfs(node *TreeNode, level int) {
	if node == nil {
		return
	}
	if len(result) < level+1 {
		result = append(result, make([]int, 0))
	}
	result[level] = append(result[level], node.Val)
	dfs(node.Left, level+1)
	dfs(node.Right, level+1)
}
