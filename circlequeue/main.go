package main

import (
	"errors"
	"fmt"
)

const queueSize = 100

type CircleQueue struct {
	data  [queueSize]interface{} // 存储数据的结构
	front int                    // 头部数据
	rear  int                    // 尾部数据
}

func main() {
	fmt.Println("circlequeue")
	var q CircleQueue
	q.InitQueue()
}

func (q *CircleQueue) InitQueue() { // 循环队列初始化
	q.front = 0
	q.rear = 0
}

func (q *CircleQueue) QueueLength() int {
	return (q.rear - q.front + queueSize) % queueSize
}

func (q *CircleQueue) EnQueue(data interface{}) error {
	if (q.rear+1)%queueSize == q.front%queueSize {
		return errors.New("Queue is full!")
	}
	q.data[q.rear] = data // Enter queue
	q.rear = (q.rear + 1) % queueSize
	return nil
}

func (q *CircleQueue) DeQueue() (data interface{}, err error) {
	if q.rear == q.front {
		return nil, errors.New("Queue is empty!")
	}
	res := q.data[q.front]              // Get the first data
	q.data[q.front] = 0                 // Clear data
	q.front = (q.front + 1) % queueSize // <100 ，>100 % queueSize
	return res, nil
}
