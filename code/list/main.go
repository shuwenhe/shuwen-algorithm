package main

import(
	"fmt"
)

type ListNode struct{
	Val int 
	Next *ListNode
}

func main()  {
	headNode := &ListNode{}
	listData := headNode

	InsertTail(1,listData,headNode)
	PrintList(listData)

	InsertTail(2,listData,headNode)
	PrintList(listData)

	InsertTail(3,listData,headNode)
	PrintList(listData)
}

func InsertTail(value int,list,position *ListNode)  {
	tempCell := new(ListNode)

	if tempCell == nil{
		fmt.Println("out of space")
	}

	tempCell.Val = value
	tempCell.Next = position.Next
	position.Next = tempCell
}

func PrintList(list *ListNode)  {
	if list.Next != nil{
		fmt.Print(list.Val,"->")
		PrintList(list.Next)
	}else{
		fmt.Println(list.Val)
	}
}
