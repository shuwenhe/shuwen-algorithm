package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val // [4,5,1,9] => [4,1,1,9]
	node.Next = node.Next.Next // [4,1,1,9] => [4,1,9] Delete third node
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
	fmt.Println("n1 = ", n1.Val)
	fmt.Println("n2 = ", n2.Val)
	fmt.Println("n3 = ", n3.Val)
	fmt.Println("n4 = ", n4.Val)
	deleteNode(n3)
	fmt.Println("***===***")
	fmt.Println("n1 = ", n1.Val)
	fmt.Println("n2 = ", n2.Val)
	fmt.Println("n3 = ", n3.Val)
	fmt.Println("n4 = ", n4.Val)
}
