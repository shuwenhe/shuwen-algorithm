// 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
// 示例：
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
// 当删除了倒数第二个节点后，链表变为 1->2->3->5.
// 说明：
// 给定的 n 保证是有效的。
// 进阶：
// 你能尝试使用一趟扫描实现吗？

package main

//定义一个链表结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	i := 0
	one, two := head, head
	for one.Next != nil {
		i++
		if i > n {
			two = two.Next
		}
		one = one.Next
	}
	//当n是倒数最大时（也就是正数第一个），i是不会大于n的
	//这其实删除的是链表的头节点
	if i < n {
		head = head.Next
		return head
	}
	//将要删除的节点的前驱节点指向要删除的节点的后
	next := two.Next
	two.Next = next.Next
	return head
}
