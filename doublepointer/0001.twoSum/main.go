// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// Example:
// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].
package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 18
	arr := twoSum(nums, target)
	fmt.Println("arr = ", arr)
}

func twoSum(nums []int, target int) []int { // 时间复杂度O(n^2)
	for i := 0; i < len(nums); i++ { // 时间复杂度O(n)
		for j := 0; j < len(nums); j++ { // 时间复杂度O(n)
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
