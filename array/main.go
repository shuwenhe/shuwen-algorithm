package main

import (
	"algorithm/array/arraylist"
	"fmt"
)

func main() {
	// list := NewArrayList()
	var list arraylist.List = arraylist.NewArrayList() // 定义接口对象，赋值的对象必须实现接口的所有方法
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append("a")
	list.Append("b")
	list.Append("c")
	for i := 0; i < 3; i++ {
		list.Insert(2, "val7")
		fmt.Println("list-insert = ", list)
	}
	list.Delete(2) // CRUD
	fmt.Println("list-delete = ", list)
}
