package main

import (
	"fmt"
)

type ListNode struct {
	val  string
	next *ListNode
}

func main() {
	head := &ListNode{}
	list1 := &ListNode{val: "张三"}
	list2 := &ListNode{val: "李四"}
	TailInsert(head, list1)
	TailInsert(head, list2)
	PrintNode(head)
}

func TailInsert(head *ListNode, newListNode *ListNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newListNode
}

func PrintNode(head *ListNode) {
	temp := head
	if temp.next == nil {
		return
	}
	for {
		fmt.Printf("%s", temp.next.val)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}
