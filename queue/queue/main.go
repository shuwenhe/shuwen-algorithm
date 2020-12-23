package main

import (
	"fmt"
)

// MyQueue interface
type MyQueue interface {
	Size() int                // Queue size
	Front() interface{}       // First data
	End() interface{}         // End data
	IsEmpty() bool            // Whether the queue is empty
	EnQueue(data interface{}) // Enter the data to queue
	DeQueue() interface{}     // Delete the data from queue
	Clear()                   // Clear the data from queue
}

// ArrayQueue struct
type ArrayQueue struct {
	dataStore []interface{} // Queue store data into the slice
	theSize   int           // Queue size
}

// NewQueue New Queue
func NewQueue() *ArrayQueue {
	q := new(ArrayQueue)
	q.dataStore = make([]interface{}, 0)
	q.theSize = 0
	return q
}

func main() {
	q := NewQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	fmt.Println("q.theSize = ", q.theSize)
	fmt.Println("queue = ", q.dataStore)
	for _, data := range q.dataStore {
		fmt.Println("q.dataStore = ", data)
	}
	fmt.Println("***")
	q.DeQueue()
	fmt.Println("q.theSize = ", q.theSize)
	fmt.Println("queue = ", q.dataStore)
	for _, data := range q.dataStore {
		fmt.Println("q.dataStore = ", data)
	}
	fmt.Println("q.IsEmpty", q.IsEmpty())
	fmt.Println("q.Front = ", q.Front())
	fmt.Println("q.End = ", q.End())
	q.Clear()
	fmt.Println("q.theSize = ", q.theSize)
}

// Size Queue size
func (q *Queue) Size() int {
	return q.theSize
}

// Front Get the first data
func (q *Queue) Front() interface{} {
	if q.theSize == 0 { // Queue is empty
		return nil // Return nil
	}
	return q.dataStore[0] // Return the first data
}

// End Get the last data of the queue
func (q *Queue) End() interface{} {
	if q.theSize == 0 {
		return nil
	}
	return q.dataStore[q.theSize-1] // 0,1,2,...,theSize-1
}

// IsEmpty Whether the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.theSize == 0
}

// EnQueue Enter the queue
func (q *Queue) EnQueue(data interface{}) {
	q.dataStore = append(q.dataStore, data)
	q.theSize++
}

// DeQueue Delete queue
func (q *Queue) DeQueue() interface{} {
	if q.theSize == 0 {
		return nil
	}
	data := q.dataStore[0]
	if q.Size() > 1 { // Queue size > 1
		q.dataStore = q.dataStore[1:q.Size()] // Get the some data of the queue
	}
	q.theSize-- // KFC 853 --
	return data // 854 855
}

// Clear Clear queue
func (q *Queue) Clear() {
	q.dataStore = make([]interface{}, 0)
	q.theSize = 0
}
