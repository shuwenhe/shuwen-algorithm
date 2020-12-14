package arrayiterator

import "algorithm/array/arraylist"

// Iterator Iterator interface
type Iterator interface {
	HasNext() bool              // 是否有下一个
	Next() (interface{}, error) // 下一个
}

// Iterable Iterable interface
type Iterable interface {
}

// ArrayIterator Array Iterator struct 构造指针访问数组
type ArrayIterator struct {
	list          *arraylist.ArrayList // 数组指针
	currentIndext int                  // Current Index
}
