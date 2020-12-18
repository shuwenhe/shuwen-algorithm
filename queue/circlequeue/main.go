package main

import (
	"errors"
	"fmt"
)

const queueSize = 10

// CircleQueue Circle queue
type CircleQueue struct {
	Data  [queueSize]interface{}
	Front int // 头部的位置
	Rear  int // 尾部的位置
}

func main() {
	fmt.Println("circleQueue")
	q := &CircleQueue{}
	InitQueue(q)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue("a")
	q.EnQueue("b")
	q.EnQueue("c")
	fmt.Println("q = ", q)

	q.DeQueue()
	fmt.Println("q = ", q)

	fmt.Println("queueLength = ", q.QueueLength())
}

// InitQueue Init queue
func InitQueue(q *CircleQueue) { // 初始化，头部尾部重合，为空
	q.Front = 0
	q.Rear = 0
}

// QueueLength Queue length
func (q *CircleQueue) QueueLength() int {
	return (q.Rear - q.Front + queueSize) % queueSize
}

// EnQueue 入队
func (q *CircleQueue) EnQueue(data interface{}) error {
	if (q.Rear+1)%queueSize == q.Front%queueSize {
		return errors.New("Queue is full!")
	}
	q.Data[q.Rear] = data             // 入队
	q.Rear = (q.Rear + 1) % queueSize // Loop 101 -> 1
	return nil
}

// DeQueue 出队
func (q *CircleQueue) DeQueue() (data interface{}, err error) {
	if q.Front == q.Rear {
		return nil, errors.New("Queue is empty!")
	}
	res := q.Data[q.Front]              // Get the first data of the queue
	q.Data[q.Front] = 0                 // Clear data
	q.Front = (q.Front + 1) % queueSize // Delete of the data index + 1
	return res, nil
}
