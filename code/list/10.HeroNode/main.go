package main

import "fmt"

// 定义单链表结点
type ListNode struct {
	val  string
	next *ListNode
}

func main() {
	// 先创建头结点
	head := &ListNode{}
	// 创建一个新结点
	person1 := &ListNode{
		val: "张三",
	}
	person2 := &ListNode{
		val: "李四",
	}
	TailInsert(head, person1)
	TailInsert(head, person2)
	PrintNode(head)
}

// 尾部插入
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

// 遍历链表
func PrintNode(head *ListNode) {
	//
	temp := head

	if temp.next == nil {
		fmt.Println("空链表")
		return
	}

	for {
		fmt.Printf("%s->", temp.next.val)
		// 是否就链表尾
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}
