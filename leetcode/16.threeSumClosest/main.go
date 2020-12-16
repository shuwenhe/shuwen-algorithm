package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println("sum = ", threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		n    = len(nums)
		best = math.MaxInt32
	)
	update := func(cur int) { // 根据差值的绝对值来更新答案
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}
	for i := 0; i < n; i++ { // 枚举 a
		if i > 0 && nums[i] == nums[i-1] { // 保证和上一次枚举的元素不相等
			continue
		}
		j, k := i+1, n-1 // 使用双指针枚举 b 和 c
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target { // 如果和为target直接返回答案
				return target
			}
			update(sum)
			if sum > target {
				k0 := k - 1                         // 如果和大于target，移动c对应的指针
				for j < k0 && nums[k0] == nums[k] { // 移动到下一个不相等的元素
					k0--
				}
				k = k0
			} else {
				j0 := j + 1                         // 如果和小于target，移动b对应的指针
				for j0 < k && nums[j0] == nums[j] { // 移动到下一个不相等的元素
					j0++
				}
				j = j0
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
