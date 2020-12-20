package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 1, 3, 3}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
