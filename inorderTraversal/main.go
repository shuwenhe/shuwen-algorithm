// 给定一个二叉树，返回它的中序 遍历。
// 示例:
// 输入: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3
// 输出: [1,3,2]
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := new(TreeNode)
	MorrisTraverMid(root)
}

func MorrisTraverMid(root *TreeNode) []int {
	// 莫里斯遍历利用叶子节点左右空域存储遍历前驱和后继
	// 达到时间复杂度O(N)，空间复杂度O(1)
	// 二叉树的串行遍历

	// 莫里斯中序遍历
	if root == nil {
		return nil
	}

	var result []int

	// 游标节点初始化为根节点
	cur := root

	// 定义前驱节点
	var pre *TreeNode

	// 当没有遍历到最后情况
	for cur != nil {

		// 当游标节点没有左孩子
		if cur.Left == nil {
			// 加游标节点值加入结果集(visit)
			result = append(result, cur.Val)

			// 因为没有左孩子，进入右孩子
			cur = cur.Right
			continue
		}

		// 当游标有左孩子
		// 1.找到左子树最右节点作为游标节点前驱

		// 得到左子树
		pre = cur.Left

		// 找到左子树最右叶子节点
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}

		// 最右叶子节点
		if pre.Right == nil {
			// 最右叶子节点与cur连接
			pre.Right = cur

			// 进入左子树
			cur = cur.Left
			continue
		}

		// 最右节点与cur相等（成环情况）
		// visit cur
		result = append(result, cur.Val)

		// 破环
		pre.Right = nil

		// 进入右子树
		cur = cur.Right
	}

	return result
}
