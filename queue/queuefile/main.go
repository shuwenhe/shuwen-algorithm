package main

import (
	"fmt"
)

// MyQueue Interface
type MyQueue interface {
	Size() int                // 队列的大小
	Front() interface{}       // 队列的第一个元素
	End() interface{}         // 队列的最后一个元素
	IsEmpty() bool            // 判断队列是否位空
	EnQueue(data interface{}) // 入队
	DeQueue() interface{}     // 出队
	Clear()                   // 清空数据
}

// Queue Struct
type Queue struct {
	dataStore []interface{} // 队列数据的存储
	theSize   int           // 队列的大小
}

func main() {
	fmt.Println("queue")
	myq := NewQueue()

	myq.EnQueue(1)
	myq.EnQueue(2)
	myq.EnQueue(3)
	myq.EnQueue("a")
	myq.EnQueue("b")
	myq.EnQueue("c")

	myq.DeQueue()

	fmt.Println("myq = ", myq)
	fmt.Println("Front = ", myq.Front())
	fmt.Println("End = ", myq.End())
	myq.Clear()
	fmt.Println("IsEmpty = ", myq.IsEmpty())
}

// NewQueue New queue
func NewQueue() *Queue {
	myQueue := new(Queue)
	myQueue.dataStore = make([]interface{}, 0)
	myQueue.theSize = 0
	return myQueue
}

// Size 队列的大小
func (myq *Queue) Size() int {
	return myq.theSize
}

// Front 队列的第一个元素
func (myq *Queue) Front() interface{} {
	if myq.theSize == 0 {
		return nil
	}
	return myq.dataStore[0]
}

// End 队列的最后一个元素
func (myq *Queue) End() interface{} {
	if myq.theSize == 0 {
		return nil
	}
	return myq.dataStore[myq.Size()-1]
}

// IsEmpty 判断队列是否位空
func (myq *Queue) IsEmpty() bool {
	return myq.theSize == 0
}

// EnQueue 入队
func (myq *Queue) EnQueue(data interface{}) {
	myq.dataStore = append(myq.dataStore, data)
	myq.theSize++
}

// DeQueue 出队
func (myq *Queue) DeQueue() interface{} {
	if myq.theSize == 0 {
		return nil
	}
	data := myq.dataStore[0]
	if myq.Size() > 1 {
		myq.dataStore = myq.dataStore[1:myq.Size()] // 对切片进行截取
	}
	myq.theSize--
	return data
}

// Clear 清空数据
func (myq *Queue) Clear() {
	myq.dataStore = make([]interface{}, 0)
	myq.theSize = 0
}
