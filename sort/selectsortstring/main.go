package main

import (
	"fmt"
)

func main() {
	arr := []string{"a", "e", "d", "c", "b"}
	fmt.Println("arr = ", arr)
	fmt.Println("***After selectsort***")
	fmt.Println("res = ", selectSortString(arr))
}

func selectSortString(arr []string) []string {
	length := len(arr)
	if length == 1 { // [1]
		return arr
	} else { // [e,d,c,b,a](Desc)
		for i := 0; i < length-1; i++ { // Last data is min
			min := i // The targer of the min
			for j := i + 1; j < length; j++ {
				if arr[j] > arr[min] {
					min = j
				}
			}
			if i != min { // Exchange the data
				arr[min], arr[i] = arr[i], arr[min]
			}
		}
		return arr
	}
}
