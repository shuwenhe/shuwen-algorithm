package main

import (
	"algorithm/array-list/ArrayList"
	"fmt"
)

func main() {
	list := ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	fmt.Println(list)
}
