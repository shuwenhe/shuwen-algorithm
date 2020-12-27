package main

import (
	"algorithm/array/1.arraylist/code/stackarray"
	"fmt"
)

func main() {
	myStack := stackarray.NewStack()
	myStack.Push(1)
	// fmt.Println(myStack.Pop())
	myStack.Push(2)
	// fmt.Println(myStack.Pop())
	myStack.Push(3)
	// fmt.Println(myStack.Pop())
	fmt.Println("myStack = ", myStack)
}
