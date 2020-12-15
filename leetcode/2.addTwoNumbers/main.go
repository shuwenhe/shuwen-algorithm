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

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 243,
	}
	l2 := &ListNode{
		Val: 564,
	}
	l3 := addTwoNumbers(l1, l2)
	fmt.Println(l3.Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	promotion := 0     // 进位值, 只可能为0或1
	var head *ListNode // 结果表的头结点
	var rear *ListNode // 保存结果表的尾结点
	for nil != l1 || nil != l2 {
		sum := 0
		if nil != l1 {
			sum += l1.Val
			l1 = l1.Next
		}
		if nil != l2 {
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

		if nil == head {
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
