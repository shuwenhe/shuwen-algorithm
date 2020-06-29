// 给定一个整数数组nums和一个目标值target，
// 请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
// 示例:
// 给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]

package main

import(
	"fmt"
)

func main()  {
	nums := []int{2,7,11,15}
	target := 26
	arr := twoSum(nums,target)
	fmt.Println("arr = ",arr)
}

func twoSum(nums []int, target int) []int {
    h := make(map[int]int) // hash table
    for i, value := range nums {
		h[target-value] = i
        if wanted, ok := h[value]; ok {
            return []int{wanted, i}
        }
    }
    return nil
}