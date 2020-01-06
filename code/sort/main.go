package main

import (
	"fmt"
)

// selection sort
func main1() {
	s := []int{2, 3, 8, 9, 0, 7, 5, 6, 1}
	fmt.Println(s)
	SelectionSort(s)
	fmt.Println(s)
}

func SelectionSort(s []int) {
	l := len(s)
	fmt.Println(l)
}

// quick sort
func main()  {
	var z []int
	quickSort(z)	
}

func quickSort(list []int)  {
	if len(list) <= 1{
		return
	}
	fmt.Println(quickSort)
}

