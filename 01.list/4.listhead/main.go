package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	headNode := &ListNode{}
	listData := headNode
	InsertTail(7, listData, headNode)
	PrintList(listData)
	InsertTail(8, listData, headNode)
	PrintList(listData)
	InsertTail(9, listData, headNode)
	PrintList(listData)
}

// 每次在链表头部插入
func InsertTail(value int, list *ListNode, position *ListNode) {
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
		fmt.Print(list.Val, "->")
		PrintList(list.Next)
	} else {
		fmt.Println(list.Val)
	}
}
