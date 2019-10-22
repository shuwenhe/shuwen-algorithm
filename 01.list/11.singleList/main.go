package main

import (
	"fmt"
)

type ListNode struct {
	name string
	next *ListNode
}

func main() {
	head := &ListNode{}
	list1 := &ListNode{
		name: "张三",
	}
	list2 := &ListNode{
		name: "李四",
	}
	TailInsert(head, list1)
	TailInsert(head, list2)
	PrintList(head)
}

func TailInsert(head *ListNode, newListNode *ListNode) {
	temp := head
	for {
		if temp.next == nil {
			// fmt.Println("空链表")
			break
		}
		temp = temp.next
	}
	temp.next = newListNode
}

func PrintList(head *ListNode) {
	temp := head
	if temp == nil {
		fmt.Println("空链表")
		return
	}
	for {
		fmt.Printf("%s", temp.next.name)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}
