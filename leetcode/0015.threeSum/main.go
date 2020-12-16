package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	arr := threeSum(nums)
	fmt.Println("arr = ", arr)
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	for first := 0; first < n; first++ { // 枚举 a
		if first > 0 && nums[first] == nums[first-1] { // 需要和上一次枚举的数不相同
			continue
		}
		third := n - 1 // c 对应的指针初始指向数组的最右端
		target := -1 * nums[first]
		for second := first + 1; second < n; second++ { // 枚举 b
			if second > first+1 && nums[second] == nums[second-1] { // 需要和上一次枚举的数不相同
				continue
			}
			for second < third && nums[second]+nums[third] > target { // 需要保证 b 的指针在 c 的指针的左侧
				third--
			}
			if second == third { // 如果指针重合，随着b后续的增加就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
