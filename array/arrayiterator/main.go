package main

import (
	"fmt"
)

// Iterator Iterator interface
type Iterator interface {
	HasNext() bool              // 是否有下一个
	Next() (interface{}, error) // 下一个
}

// Iterable Iterable interface
type Iterable interface {
}

// ArrayList struct
type ArrayList struct {
	dataStore []interface{} // Array 任意类型的
	TheSize   int           // Array的大小
}

// ArrayIterator Array Iterator struct 构造指针访问数组
type ArrayIterator struct {
	list          *ArrayList // 数组指针
	currentIndext int        // Current Index
}

func main() {
	fmt.Println("arrayiterator")
}
