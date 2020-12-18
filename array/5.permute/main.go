// 给定一个没有重复数字的序列，返回其所有可能的全排列。
// 示例:
// 输入: [1,2,3]
// 输出:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	permutes := permute(nums)
	fmt.Println(permutes)
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	res := [][]int{}
	for i := range nums {
		l := make([]int, len(nums)-1)
		copy(l[:i], nums[:i])
		copy(l[i:], nums[i+1:])
		for _, x := range permute(l) {
			//以 nums[i] 为首的排列
			res = append(res, append([]int{nums[i]}, x...))
		}
	}
	return res
}
