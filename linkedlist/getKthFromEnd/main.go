// 实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
// 注意：本题相对原题稍作改动
// 示例：
// 输入： 1->2->3->4->5 和 k = 2
// 输出： 4
// 说明：
// 给定的 k 保证是有效的。
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("kthToLast")
	n5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	n4 := &ListNode{
		Val:  4,
		Next: n5,
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
	fmt.Println(getKthFromEnd(n1, 2))
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast := head
	slow := head
	for k > 0 && fast != nil{
		fast = fast.Next
		k--
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
