// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// Example:
// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].
package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15} // Given an integer array nums
	target := 18                // Target value
	arr := twoSum(nums, target) // Find the two integers whose sum is the target value in the array, and return their array index
	fmt.Println("arr = ", arr)  // Print subscript
}

func twoSum(nums []int, target int) []int { // The time complexity of hash lookup is O(1)
	hashTable := map[int]int{} // Hash container map reduces time complexity
	for i, x := range nums {   // Enumerate each number x in the array
		if p, ok := hashTable[target-x]; ok { // Find target-x
			return []int{p, i}
		}
		hashTable[x] = i // Insert x into the hash table to ensure that x will not match itself.
	}
	return nil
}
