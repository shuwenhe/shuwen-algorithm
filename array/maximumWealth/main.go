package main

import "fmt"

func main() {
	accounts := [][]int{{1, 5}, {7, 3}, {6, 5}}
	fmt.Println(maximumWealth(accounts))
}

func maximumWealth(accounts [][]int) (res int) {
	for _, account := range accounts {
		sum := 0
		for _, val := range account {
			sum += val
		}
		if res < sum {
			res = sum
		}
	}
	return
}
