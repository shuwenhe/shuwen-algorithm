package main

type Node struct {
	data  interface{}
	pNext *Node
}

type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() data interface{}
}
