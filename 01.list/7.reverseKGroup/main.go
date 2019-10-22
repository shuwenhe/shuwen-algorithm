// 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
// 示例 :
// 给定这个链表：1->2->3->4->5
// 当 k = 2 时，应当返回: 2->1->4->3->5
// 当 k = 3 时，应当返回: 3->2->1->4->5
// 说明 :
// 你的算法只能使用常数的额外空间。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// reverseKGroup()
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 || head == nil || head.Next == nil {
		return head
	}

	tail, needReverse := getTail(head, k)

	if needReverse {
		tailNext := tail.Next
		//斩断 tail 后的链接
		tail.Next = nil
		head, tail = reverse(head, tail)

		//tail 后面接上尾部的递归处理
		tail.Next = reverseKGroup(tailNext, k)
	}

	return head
}

func getTail(head *ListNode, k int) (*ListNode, bool) {
	for k > 1 && head != nil {
		head = head.Next
		k--
	}
	return head, k == 1 && head != nil
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	//反转链表
	prev, cur := head, head.Next
	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}
	return tail, head
}
