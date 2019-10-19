package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体管理队列
type Queue struct {
	maxSize int
	array   [5]int //数组->模拟队列
	front   int    //表示指向队列首部
	rear    int    //表示指向队列尾部

}

//添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {
	//先判断队列是否已满
	if this.rear == this.maxSize-1 { //重要提示：rear是队列尾部（含最后元素）
		return errors.New("queue full")
	}
	this.rear++ //rear 后移
	this.array[this.rear] = val
	return
}

//从队列中取数据
func (this *Queue) GetQueue() (val int, err error) {
	//先判断队列是否为空
	if this.rear == this.front { //队空
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val, err
}

//显示队列，找到队首，然后遍历到队尾。
func (this *Queue) showQueue() {
	fmt.Println("队列当前的情况是：")
	//this.front不包含队首元素。
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
}

//编写一个主函数，测试。
func main() {
	//先创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
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
			err := queue.AddQueue(val)
			if err != nil {

				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.showQueue()
		case "exit":
			os.Exit(0)
		}
	}
	//上面的代码实现了基本队列结构，但是没有有效的利用数组空间。
	//请思考，如何使用数组实现一个环形的队列。
}
