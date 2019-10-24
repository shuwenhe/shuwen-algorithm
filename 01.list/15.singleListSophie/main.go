package main

import "fmt"

// 声明链表
type ListNode struct {
	Name string
	Next *ListNode // 递归声明，多个Node
}

func main() {
	HeadNode := &ListNode{}
	NewNodeRichard := &ListNode{Name: "Richard"}
	NewNodeRitchie := &ListNode{Name: "Ritchie"}
	InsertNode(HeadNode, NewNodeRichard)
	InsertNode(HeadNode, NewNodeRitchie)
	QueryNode(HeadNode)
}

// CRUD(Create,Read,Update,Delete)
// TailInsert HeadNode(temp)->Richard(temp)->Ritchie->...->noden(Tail)->NewNode
// HeadInsert
func InsertNode(HeadNode *ListNode, NewNode *ListNode) {
	temp := HeadNode
	// 找尾巴
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	temp.Next = NewNode
}

// HeadNode(temp)->Richard(temp)->Ritchie(temp)->...->noden(Tail)->NewNode
func QueryNode(HeadNode *ListNode) {
	// temp := HeadNode
	// if temp.Next == nil {
	if HeadNode.Next == nil {
		return
	}
	// for {
	// 	if temp.Next != nil {
	// 		fmt.Println(temp.Next.Name)
	// 		temp = temp.Next
	// 	} else {
	// 		return
	// 	}
	// }
	for {
		// fmt.Println(temp.Next.Name)
		fmt.Println(HeadNode.Next.Name)
		HeadNode = HeadNode.Next
		if HeadNode.Next == nil {
			return
		}
	}
}
