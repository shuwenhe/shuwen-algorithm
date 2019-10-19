package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 8, 9, 0, 7, 5, 6, 1}
	fmt.Println(s)
	SelectionSort(s)
	fmt.Println(s)
}

func SelectionSort(s []int) {
	l := len(s)
	fmt.Println(l)
}
