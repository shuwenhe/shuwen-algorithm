package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	headNode *ListNode // head node
}

func main() {
	headNode := &ListNode{}
	listData := headNode
	Insert(1, listData, headNode)
	Insert(2, listData, headNode)
	Insert(3, listData, headNode)
	PrintList(listData)
	// fmt.Println(listData)
}

// 1.Insert
func Insert(value int, list *ListNode, position *ListNode) {
	tempCell := new(ListNode)
	if tempCell == nil {
		fmt.Println("out of space")
	}
	tempCell.Val = value
	tempCell.Next = position.Next
	position.Next = tempCell
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
