// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
// 示例 1:
// 输入: "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:
// 输入: "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
// 输入: "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

package main

import (
	"fmt"
	"strings"
)

func main() {
	i := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(i)
}
func lengthOfLongestSubstring(s string) int {
	var max, start, end, length = 0, 0, 0, len(s)
	for start < length && end < length {
		if !strings.Contains(s[start:end], string(s[end])) {
			end = end + 1
			if end-start > max {
				max = end - start
			}
		} else {
			start = start + 1
		}
	}
	return max
}
