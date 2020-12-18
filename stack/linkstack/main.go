package main

import (
	"fmt"
)

// type LinkStack interface {
// 	IsEmpty() bool
// 	Push(data interface{})
// 	Pop() interface{}
// 	Length() int
// }

type Node struct {
	Data  interface{}
	PNext *Node
}

func NewStack() *Node {
	return &Node{}
}

func main() {
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

func (n *Node) Push(data interface{}) {
}

func (n *Node) Pop() interface{} {
	return nil
}

func (n *Node) Length() int {
	return 0
}
