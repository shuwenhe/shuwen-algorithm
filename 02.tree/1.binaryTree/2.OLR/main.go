package main

import (
	"fmt"
)

type Node struct {
	Name string
	L, R *Node
}

func (node *Node) Select() {
	fmt.Print(node.Name, " ")
}

func (node *Node) SetValue(v string) {
	if node == nil {
		fmt.Println("setting value to nil.node ignored.")
		return
	}
	node.Name = v
}

func (node *Node) OLR() {
	if node == nil {
		return
	}
	node.Select()
	node.L.OLR()
	node.R.OLR()
}

func CreateNode(v string) *Node {
	return &Node{Name: v}
}

func main() {
	root := Node{Name: "ShuwenHe"}
	root.L = &Node{}
	root.L.SetValue("Ritchie")
	root.L.R = CreateNode("RitchieSon")
	root.R = &Node{"Richard", nil, nil}
	root.R.L = CreateNode("RichardSon")

	fmt.Print("\nOLR: ")
	root.OLR()
}
