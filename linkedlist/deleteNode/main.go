package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) int {
	return node.Val
}

func main() {
	fmt.Println("deleteNode")
	n4 := &ListNode{
		Val:  4,
		Next: nil,
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
	fmt.Println("i = ", deleteNode(n1))
	fmt.Println("i = ", deleteNode(n2))
	fmt.Println("i = ", deleteNode(n3))
	fmt.Println("i = ", deleteNode(n4))
	fmt.Println("i = ", deleteNode(n1.Next))
}
