// 给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
// 示例:
// 输入: 3
// 输出: 5
// 解释:
// 给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
//    1         3     3      2      1
//     \       /     /      / \      \
//      3     2     1      1   3      2
//     /     /       \                 \
//    2     1         2                 3

package main

import (
	"fmt"
)

func main() {
	nums := numTrees(3)
	fmt.Println(nums)
}

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}
	return dp[n]
}