package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	vals := []int{}
	for ; head != nil; head = head.Next { // 复制链表值到数组列表中
		vals = append(vals, head.Val)
	}
	n := len(vals)
	for i, v := range vals[:n/2] {
		if v != vals[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("isPalindrome")
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	fmt.Println(isPalindrome(head))
	fmt.Println(head.Val)
}
