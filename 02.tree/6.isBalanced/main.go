// 给定一个二叉树，判断它是否是高度平衡的二叉树。
// 本题中，一棵高度平衡二叉树定义为：
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
// 示例 1:
// 给定二叉树 [3,9,20,null,null,15,7]
//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回 true 。
// 示例 2:
// 给定二叉树 [1,2,2,3,3,null,null,4,4]
//        1
//       / \
//      2   2
//     / \
//    3   3
//   / \
//  4   4
// 返回 false 。

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isBalanced(root.Left) && isBalanced(root.Right) && abs(maxDepth(root.Left)-maxDepth(root.Right)) <= 1
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
