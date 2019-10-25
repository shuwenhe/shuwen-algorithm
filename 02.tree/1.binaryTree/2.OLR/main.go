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
	// 判断子树根结点是否为空
	if RootNode == nil {
		return
	}
	// 读取子树RootNode value
	RootNode.ReadNode()
	RootNode.Left.OLR()
	RootNode.Right.OLR()
}

// 每次读取子树的根结点
func (RootNode *Node) ReadNode() {
	fmt.Println(RootNode.Value)
}
