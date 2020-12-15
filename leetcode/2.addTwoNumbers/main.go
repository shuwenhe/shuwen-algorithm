package main

import (
	"fmt"
)

// 给出两个非空的链表用来表示两个非负的整数。其中，它们各自的位数是按照逆序的方式存储的，
// 并且它们的每个节点只能存储一位数字。
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
// 您可以假设除了数字0之外，这两个数都不会以0开头。
// 示例：
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807

// ListNode Struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1Node3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	list1Node2 := &ListNode{
		Val:  4,
		Next: list1Node3,
	}
	list1Node1 := &ListNode{
		Val:  2,
		Next: list1Node2,
	}

	list2Node3 := &ListNode{
		Val:  4,
		Next: nil,
	}
	list2Node2 := &ListNode{
		Val:  6,
		Next: list2Node3,
	}
	list2Node1 := &ListNode{
		Val:  5,
		Next: list2Node2,
	}

	res := addTwoNumbers(list1Node1, list2Node1)
	for {
		fmt.Println(res.Val)
		if res.Next == nil {
			break
		}
		res = res.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	promotion := 0               // Carry value, can only be 0 or 1
	var head *ListNode           // Head node of the result table
	var rear *ListNode           // Rear node of the result table
	for l1 != nil || l2 != nil { // Traverse two linked lists at the same time
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += promotion
		promotion = 0

		if sum >= 10 {
			promotion = 1
			sum = sum % 10
		}

		node := &ListNode{
			sum,
			nil,
		}

		if head == nil {
			head = node
			rear = node
		} else {
			rear.Next = node
			rear = node
		}
	}

	if promotion > 0 {
		rear.Next = &ListNode{
			promotion,
			nil,
		}
	}
	return head
}
