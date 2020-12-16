package main

import (
	"fmt"
)

func main() {
	strs := []string{"flower", "flow", "flight", "feat"}
	str := longestCommonPrefix(strs)
	fmt.Println("str = ", str)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 { // 如果字符数组为空
		return ""
	}
	var lcp func(int, int) string
	lcp = func(start, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2
		lcpLeft, lcpRight := lcp(start, mid), lcp(mid+1, end)
		minLength := min(len(lcpLeft), len(lcpRight))
		for i := 0; i < minLength; i++ {
			if lcpLeft[i] != lcpRight[i] {
				return lcpLeft[:i]
			}
		}
		return lcpLeft[:minLength]
	}
	return lcp(0, len(strs)-1)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
