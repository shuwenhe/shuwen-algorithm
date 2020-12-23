package main

import (
	"fmt"
)

// LinkStack interface
type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}

// Node struct
type Node struct {
	Data  interface{} // Data of the stack
	PNext *Node       // Next pointer
}

// NewStack New stack
func NewStack() *Node {
	return &Node{}
}

func main() {
	myStack := NewStack()
	count := 10
	for i := 0; i < count; i++ {
		myStack.Push(i)
	}
	for data := myStack.Pop(); data != nil; data = myStack.Pop() {
		fmt.Println(data)
	}
}

func main2() {
	n1 := new(Node)
	n2 := new(Node)
	n3 := new(Node)
	n4 := new(Node)
	n1.Data = 1
	n1.PNext = n2
	n2.Data = 2
	n2.PNext = n3
	n3.Data = 3
	n3.PNext = n4
	n4.Data = 4
	fmt.Println("n1 = ", n1)
	fmt.Println("n1.data = ", n1.Data)
	fmt.Println("n1.PNext = ", n1.PNext)
	fmt.Println("****")
	fmt.Println("n2 = ", n2)
	fmt.Println("n2.data = ", n2.Data)
	fmt.Println("n2.PNext = ", n2.PNext)
	fmt.Println("****")
	fmt.Println("n3 = ", n3)
	fmt.Println("n3.data = ", n3.Data)
	fmt.Println("n3.PNext = ", n3.PNext)
	fmt.Println("****")
	fmt.Println("n4 = ", n4)
	fmt.Println("n4.data = ", n4.Data)
	fmt.Println("n4.PNext = ", n4.PNext)
}

// IsEmpty bool
func (n *Node) IsEmpty() bool {
	if n.PNext == nil {
		return true
	}
	return false
}

// Push Push data into stack
func (n *Node) Push(data interface{}) {
	newNode := &Node{
		Data:  data,
		PNext: nil,
	}
	newNode.PNext = n.PNext
	n.PNext = newNode // 头部插入
}

// Pop Pop data from stack
func (n *Node) Pop() interface{} {
	if n.IsEmpty() == true {
		return nil
	}
	value := n.PNext.Data   // The data of the pop
	n.PNext = n.PNext.PNext // Delete
	return value
}

// Length The Length of the stack
func (n *Node) Length() int {
	pNext := n
	length := 0
	for pNext.PNext != nil {
		pNext = pNext.PNext
		length++
	}
	return length
}
