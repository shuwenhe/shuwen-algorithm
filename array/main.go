// 给你一个整数数组 nums 。
// 如果一组数字 (i,j) 满足 nums[i] == nums[j] 且 i < j ，就可以认为这是一组 好数对 。
// 返回好数对的数目。
// 示例 1：
// 输入：nums = [1,2,3,1,1,3]
// 输出：4
// 解释：有 4 组好数对，分别是 (0,3), (0,4), (3,4), (2,5) ，下标从 0 开始
// 示例 2：
// 输入：nums = [1,1,1,1]
// 输出：6
// 解释：数组中的每组数字都是好数对
// 示例 3：
// 输入：nums = [1,2,3]
// 输出：0
// 提示：
// 1 <= nums.length <= 100
// 1 <= nums[i] <= 100

package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	var index1, index2 int
	var r int
	for index1 < len(nums) && index2 < len(nums) {
		if index1 < index2 {
			if nums[index1] == nums[index2] {
				r++ // 双指针遍历数组，符合条件结果+1
			}
		}
		index2++                 // 每次循环将2指针+1
		if index2 == len(nums) { // 判断2指针边界，超出边界
			index1++            //  将1指针进位
			index2 = index1 + 1 // 重置2指针为1指针的+1
		}
	}
	return r
}

func main() {
	nums := []int{1, 2, 3, 1, 1, 3}
	fmt.Println(numIdenticalPairs(nums))
}
