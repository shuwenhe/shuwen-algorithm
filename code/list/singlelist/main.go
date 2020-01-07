package main

import (
	"fmt"
)

type listNode struct {
	name string
	next *listNode
}

func main() {
	headNode := &listNode{}
	newNode1 := &listNode{name: "Richard"}
	newNode2 := &listNode{name: "Ritchie"}
	insertNode(headNode, newNode1)
	insertNode(headNode, newNode2)
	// select
	selectNode(headNode)
}

// insert tail
func insertNode(headNode *listNode, newNode *listNode) {
	temp := headNode
	// find insert space
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newNode
}

// select
func selectNode(headNode *listNode) {
	temp := headNode
	for {
		fmt.Printf("%s", temp.next.name)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}
