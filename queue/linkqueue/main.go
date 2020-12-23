package main

import (
	"fmt"
)

type Node struct {
	Data  interface{}
	PNext *Node
}

type LinkQueue struct {
	Front *Node
	Rear  *Node
}

type LinkQueueInf interface {
	Length() int
	EnQueue(data interface{})
	DeQueue() interface{}
}

func NewLinkQueue() *LinkQueue {
	return &LinkQueue{}
}

func main() {
	fmt.Println("linkqueue")
	lq := NewLinkQueue()
	for i := 0; i < 100; i++ {
		lq.EnQueue(i)
	}
	for data := lq.DeQueue(); data != nil; data = lq.DeQueue() {
		fmt.Println("data = ", data)
	}
	lq.EnQueue("a")
	fmt.Println("data = ", lq.Front.Data)
}

func (lq *LinkQueue) Length() int {
	pnext := lq.Front
	length := 0
	for pnext.PNext != nil {
		pnext = pnext.PNext
		length++
	}
	return length
}

func (lq *LinkQueue) EnQueue(data interface{}) {
	newNode := &Node{
		Data:  data,
		PNext: nil,
	}
	if lq.Front == nil {
		lq.Front = newNode
		lq.Rear = newNode
	} else {
		lq.Rear.PNext = newNode
		lq.Rear = lq.Rear.PNext
	}
}

func (lq *LinkQueue) DeQueue() interface{} {
	if lq.Front == nil {
		return nil
	}
	newNode := lq.Front
	if lq.Front == lq.Rear {
		lq.Front = nil
		lq.Rear = nil
	} else { // 3>->2->1 = 3->2
		lq.Front = lq.Front.PNext // Delete queue data
	}
	return newNode.Data
}
