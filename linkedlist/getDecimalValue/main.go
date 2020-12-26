package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	res := 0
	for tmp := head; tmp != nil; {
		res = (res << 1) | (tmp.Val)
		tmp = tmp.Next
	}
	return res
}

func main() {
	n3 := &ListNode{
		Val:  1,
		Next: nil,
	}
	n2 := &ListNode{
		Val:  0,
		Next: n3,
	}
	n1 := &ListNode{
		Val:  1,
		Next: n2,
	}
	fmt.Println(getDecimalValue(n1))
}
