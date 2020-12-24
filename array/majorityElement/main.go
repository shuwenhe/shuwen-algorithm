package main

import (
	"fmt"
)

func main() {
	fmt.Println("majorityElement")
	nums := []int{1, 2, 5, 9, 5, 9, 5, 5, 5}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	vote := 0
	for _, num := range nums {
				fmt.Println("num = ", num)
			}
	return vote
}
