// 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
// 示例 1:
// 输入: num1 = "2", num2 = "3"
// 输出: "6"
// 示例 2:
// 输入: num1 = "123", num2 = "456"
// 输出: "56088"
// 说明：
// num1 和 num2 的长度小于110。
// num1 和 num2 只包含数字 0-9。
// num1 和 num2 均不以零开头，除非是数字 0 本身。
// 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

package main

import (
	"fmt"
)

func main() {
	num1 := "24"
	num2 := "5"
	mutiply := multiply(num1, num2)
	fmt.Println(mutiply)
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	//两数相乘后，长度最长即 len(num1)+len(num2)，例如 999*999 = 998001
	temp := make([]int, len(num1)+len(num2))
	for i, a := range num1 {
		for j, b := range num2 {
			//每个位上是ASCII中的数字大小，需要-‘0’得到十进制的大小，例如 ‘0’ -> 十进制48
			temp[i+j+1] += int((a - '0') * (b - '0'))
		}
	}

	//从右往左，处理进位,首位不需要处理
	for i := len(temp) - 1; i > 0; i-- {
		temp[i-1] += temp[i] / 10
		temp[i] = temp[i] % 10
	}

	//首位可能是0，比如 12*12 = 0144
	if temp[0] == 0 {
		temp = temp[1:]
	}

	//转string
	//temp中选用 []int,而不是[]byte，
	// 是因为 byte 相当于 uint8，最大值为255
	// 考虑进位的话,temp可能会溢出
	// 例如 999*999 会得到 [0] [81] [81+81] [81+81+81] [81+81] [81]
	// 得到 0 81 162 243 162 81
	// 当进位时，243这一位会超过255，溢出
	res := make([]byte, len(temp))
	for i := 0; i < len(temp); i++ {
		res[i] = '0' + byte(temp[i])
	}

	return string(res)
}
