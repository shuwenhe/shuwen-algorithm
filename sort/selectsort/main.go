package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	fmt.Println("arr = ", arr)
	fmt.Println("***After selectsort***")
	fmt.Println("res = ", selectSort(arr))
}

// [1,3,5,7,9,2,4,6,8,0]
// [9  1,3,5,7,2,4,6,8,0]
// [9,8  1,3,5,7,2,4,6,0]
// [9,8,7  1,3,5,2,4,6,0]
// [...]
func selectSort(arr []int) []int {
	length := len(arr)
	if length == 1 { // [1]
		return arr
	} else { // [9,8,7,6,5,4,3,2,1](Desc)
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
