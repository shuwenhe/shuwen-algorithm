package main

import (
	"fmt"
)

func main() {
	i := fib(7)
	fmt.Println("i = ", i)
}

// fib 1,1,2,3,5,8,13...
func fib(num int) int {
	if num == 1 || num == 2 {
		return 1
	}
	return fib(num-1) + fib(num-2)
}
