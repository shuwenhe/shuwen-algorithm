// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// Example:
// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].
package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15} // 给定一个整数数组nums
	target := 18                // 目标值target
	arr := twoSum(nums, target) // 在该数组中找出和为目标值的那两个整数，并返回他们的数组下标
	fmt.Println("arr = ", arr)  // 打印下标
}

func twoSum(nums []int, target int) []int { // 哈希查找的时间复杂度为 O(1)
	hashTable := map[int]int{} // 哈希容器map降低时间复杂度
	for i, x := range nums {   // 枚举数组中的每一个数x
		if p, ok := hashTable[target-x]; ok { // 寻找target - x
			return []int{p, i}
		}
		hashTable[x] = i // 将x插入到哈希表中，保证不会让x和自己匹配。
	}
	return nil
}
