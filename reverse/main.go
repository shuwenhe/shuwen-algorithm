// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
// 示例 1:
// 输入: 123
// 输出: 321
// 示例 2:
// 输入: -123
// 输出: -321
// 示例 3:
// 输入: 120
// 输出: 21
// 注意:
// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

package main

import (
	"fmt"
)

func main() {
	num := reverse(123)
	fmt.Println(num)
}

// An highlighted block
func reverse(x int) int {
	var nums, newnums int
	for x != 0 { //直到x等于0，跳出循环
		a := x % 10
		newnums = nums*10 + a
		nums = newnums
		x = x / 10
		//题目要求其数值范围是 [−2^31,  2^31 − 1]。如果反转后的整数溢出，则返回 0。
		MaxInt32 := 1<<31 - 1
		MinInt32 := -1 << 31
		if nums > MaxInt32 || nums < MinInt32 {
			return 0
		}
	}
	return nums
}
