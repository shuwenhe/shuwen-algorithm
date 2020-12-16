// Given a string s, find the length of the longest substring without repeating characters.

// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// Example 2:
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

// Example 3:
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

// Example 4:
// Input: s = ""
// Output: 0
// Constraints:
// 0 <= s.length <= 5 * 104
// s consists of English letters, digits, symbols and spaces.
package main

import (
	"fmt"
)

func main() {
	s := "abcabcbb"
	i := lengthOfLongestSubstring(s)
	fmt.Println("i = ", i)
}

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{} // Hash collection, record whether each character has appeared
	n := len(s)
	rk, ans := -1, 0 // Right pointer, the initial value is -1, which is equivalent to that we are on the left side of the left boundary of the string and have not yet begun to move
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1]) // The left pointer moves one space to the right to remove a character
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++ // 不断地移动右指针
			rk++
		}
		ans = max(ans, rk-i+1) // 第 i 到 rk 个字符是一个极长的无重复字符子串
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
