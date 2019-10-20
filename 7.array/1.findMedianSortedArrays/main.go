// 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
// 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
// 你可以假设 nums1 和 nums2 不会同时为空。
// 示例 1:
// nums1 = [1, 3]
// nums2 = [2]
// 则中位数是 2.0
// 示例 2:
// nums1 = [1, 2]
// nums2 = [3, 4]
// 则中位数是 (2 + 3)/2 = 2.5

package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	medianNum := 0.0
	medianNum = findMedianSortedArrays(nums1, nums2)
	fmt.Println(medianNum)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l == 0 {
		panic("切片的长度为0，无法求解中位数。")
	}

	res := make([]int, 0, l)
	//归并排序
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	if i >= len(nums1) {
		res = append(res, nums2[j:]...)
	}
	if j >= len(nums2) {
		res = append(res, nums1[i:]...)
	}

	if l%2 == 0 {
		return float64(res[l/2]+res[l/2-1]) / 2.0
	}
	return float64(res[l/2])
}
