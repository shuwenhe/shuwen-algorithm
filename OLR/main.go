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
	root.Left = &Node{"Ritchie", nil, nil} // n层如何遍历
	root.Left.Right = &Node{"Ritchie son", nil, nil}
	root.Right = &Node{"Richard", nil, nil}
	root.Right.Left = &Node{"Richard son", nil, nil}
	root.OLR()
}

func (RootNode *Node) OLR() {
	if RootNode == nil { // 判断根结点是否为空
		return
	}

	fmt.Println(RootNode.Value) // read root node value
	RootNode.Left.OLR()
	RootNode.Right.OLR()
}

