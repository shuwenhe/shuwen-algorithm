package main

import "fmt"

// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
// 示例:
// 给定 1->2->3->4, 你应该返回 2->1->4->3.

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{}
	listData := head
	Insert(1, listData, head)
	Insert(2, listData, head)
	Insert(3, listData, head)
	Insert(4, listData, head)
	swapPairs(head)
	PrintList(listData)
	// swapPairs(listData)
	// PrintList(head)
}

func PrintList(list *ListNode) {
	if list.Next != nil {
		fmt.Println(list.Val)
		PrintList(list.Next)
	} else {
		fmt.Println(list.Val)
	}
}

func Insert(value int, list *ListNode, position *ListNode) {
	tempCell := new(ListNode)
	if tempCell == nil {
		fmt.Println("out of space")
	}
	tempCell.Val = value
	tempCell.Next = position.Next
	position.Next = tempCell
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	cur := head
	head = cur.Next
	for ; cur != nil && cur.Next != nil; cur = cur.Next {
		next := cur.Next
		//注意：第一次循环时，prev为nil
		if prev != nil {
			prev.Next = next
		}
		//交换两个节点
		cur.Next, next.Next, prev = next.Next, cur, cur
	}

	return head
}
