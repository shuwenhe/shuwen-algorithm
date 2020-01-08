package main

import (
	"fmt"
)

// struct
type ListNode struct {
	Name string
	Next *ListNode
}

func main() {
	HeadNode := &ListNode{}
	NewNodeRichard := &ListNode{Name: "Richard"}
	NewNodeRitchie := &ListNode{Name: "Ritchie"}
	InsertNode(HeadNode, NewNodeRichard)
	InsertNode(HeadNode, NewNodeRitchie)
	SelectNode(HeadNode)
}

// insert node
func InsertNode(HeadNode *ListNode, NewNode *ListNode) {
	// temp
	temp := HeadNode
	// tail space
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	temp.Next = NewNode
}

// select node
func SelectNode(HeadNode *ListNode) {
	temp := HeadNode
	for {
		fmt.Println(temp.Next.Name)
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
}
