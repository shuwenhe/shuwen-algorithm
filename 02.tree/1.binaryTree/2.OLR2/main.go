package main

import (
	"fmt"
)

type Node struct {
	Value       string
	Left, Right *Node
}

func main() {
	root := Node{Value: "Shuwen"}
	root.Left = &Node{}
	root.Left.InsertValue("Ritchie")
	root.Left.Right = CreateNode("Ritchie Son")
	root.Right = &Node{"Richard", nil, nil}
	root.Right.Left = CreateNode("Richard Son")
	fmt.Print("\nOLR:")
	root.OLR()
}

func CreateNode(value string) *Node {
	return &Node{Value: value}
}

func (node *Node) InsertValue(value string) {
	if node == nil {
		fmt.Println("Insert value to nil.node ignored.")
		return
	}
	node.Value = value
}

func (node *Node) OLR() {
	if node == nil {
		return
	}
	node.SelectNode()
	node.Left.OLR()
	node.Right.OLR()
}

func (node *Node) SelectNode() {
	fmt.Println(node.Value, "")
}
