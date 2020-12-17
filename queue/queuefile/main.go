package main

import (
	"fmt"
)

// MyQueue Interface
type MyQueue interface {
	GetSize() int             // 队列的大小
	Front() interface{}       // 队列的第一个元素
	End() interface{}         // 队列的最后一个元素
	IsEmpty() bool            // 判断队列是否位空
	EnQueue(data interface{}) // 入队
	DeQueue() interface{}     // 出队
	Clear()                   // 清空数据
}

// Queue Struct
type Queue struct {
	Data []interface{} // 队列数据的存储
	Size int           // 队列的大小
}

func main() {
	fmt.Println("queue")
	myq := Queue{}

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

// Size 队列的大小
func (myq *Queue) GetSize() int {
	return myq.Size
}

// Front 队列的第一个元素
func (myq *Queue) Front() interface{} {
	if myq.Size == 0 {
		return nil
	}
	return myq.Data[0]
}

// End 队列的最后一个元素
func (myq *Queue) End() interface{} {
	if myq.Size == 0 {
		return nil
	}
	return myq.Data[myq.GetSize()-1]
}

// IsEmpty 判断队列是否位空
func (myq *Queue) IsEmpty() bool {
	return myq.Size == 0
}

// EnQueue 入队
func (myq *Queue) EnQueue(data interface{}) {
	myq.Data = append(myq.Data, data)
	myq.Size++
}

// DeQueue 出队
func (myq *Queue) DeQueue() interface{} {
	if myq.Size == 0 {
		return nil
	}
	data := myq.Data[0]
	if myq.GetSize() > 1 {
		myq.Data = myq.Data[1:myq.GetSize()] // 对切片进行截取
	}
	myq.Size--
	return data
}

// Clear 清空数据
func (myq *Queue) Clear() {
	myq.Data = make([]interface{}, 0)
	myq.Size = 0
}
