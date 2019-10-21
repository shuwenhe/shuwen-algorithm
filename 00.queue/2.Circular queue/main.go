package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体管理环形队列
type CircleQueue struct {
	maxSize int
	array   [5]int //数组
	head    int    //指向队首
	tail    int    //指向队尾
}

//入队列AddQueue(push),GetQueue(pop)
//push是入队列
func (this *CircleQueue) push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue is full")
	}
	this.array[this.tail] = val //把值给尾部
	this.tail = (this.tail + 1) % this.maxSize
	return
}

//pop是出队列
func (this *CircleQueue) pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue is Empty")
	}
	//取出元素，head是指向队首，并且含队首元素
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return

}

//显示队列
func (this *CircleQueue) listQueue() {
	fmt.Println("环形队列情况如下")
	//取出当前队列有多少个元素
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	//设计一个辅助变量，指向head.
	tempHead := this.head
	for i := this.head; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
}

//判断环形队列是否满了。
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

//判断环形队列是否为空。
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

//取出环形队列有多少个元素
func (this *CircleQueue) Size() int {
	//这是一个关键算法
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func main() {
	//初始化一个环形队列
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示退出队列")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列的数据")
			fmt.Scanln(&val)
			err := queue.push(val)
			if err != nil {

				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.listQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
