package main

import (
	"algorithm/1.arraylist/code/arraylist"
	"fmt"
)

func main1() {
	list := arraylist.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	fmt.Println("list = ", list)
}

func main2() {
	list := arraylist.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	fmt.Println("TheSize = ", list.TheSize) // 小写是私有只能内部用，大写public
}

func main4() {
	list := arraylist.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	fmt.Println("TheSize = ", list.TheSize) // 小写是私有只能内部用，大写public
}

func main5() {
	list := arraylist.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	for i := 0; i < 4; i++ {
		list.Insert(1, "x5")
		fmt.Println("list = ", list)
	}
	fmt.Println("delete list = ", list)
	list.Delete(5)               // 删除
	fmt.Println("list = ", list) // 小写是私有只能内部用，大写public
}

func main6() {
	list := arraylist.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	for i := 0; i < 55; i++ {
		list.Insert(1, "x5")
		fmt.Println("list = ", list)
	}
	fmt.Println("list = ", list) // 小写是私有只能内部用，大写public
}

func main() {
	list := arraylist.NewArrayList()
	list.Append("a")
	list.Append("b")
	list.Append("c")
	list.Append("d")
	list.Append("e")
	for it := list.Iterator(); it.HasNext(); {
		item, _ := it.Next("111111")
		if item == "d" {
			it.Remove()
		}
		fmt.Println(item)
	}
}
