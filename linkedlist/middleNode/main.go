package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var res []*ListNode
	for head != nil {
		res = append(res, head)
		head = head.Next
	}
	return res[len(res)/2]
}

func main() {
	n5 := &ListNode{
		Val:  5,
		Next: nil,
	}

	n4 := &ListNode{
		Val:  4,
		Next: n5,
	}

	n3 := &ListNode{
		Val:  3,
		Next: n4,
	}

	n2 := &ListNode{
		Val:  2,
		Next: n3,
	}

	n1 := &ListNode{
		Val:  1,
		Next: n2,
	}
	fmt.Println(middleNode(n1).Val)
}
