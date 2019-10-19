package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 2.Print
func PrintList(list *ListNode) {
	if list.Next != nil {
		fmt.Println(list.Val)
		PrintList(list.Next)
	} else {
		fmt.Println(list.Val)
	}
}
