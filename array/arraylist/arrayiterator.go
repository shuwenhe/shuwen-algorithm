package arraylist

import (
	"errors"
)

// Iterator Iterator interface
type Iterator interface {
	HasNext() bool                             // 是否有下一个
	Next(password string) (interface{}, error) // 下一个
	Remove()                                   // Delete
	GetIndex() int                             // Get index
}

// Iterable Iterable interface
type Iterable interface {
	Iterator() Iterator // 构造初始化接口
}

// ArrayListIterator Array Iterator struct 构造指针访问数组
type ArrayListIterator struct {
	list         *ArrayList // 数组指针
	currentIndex int        // Current Index
}

// Iterator Iterator
func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator) //
	it.currentIndex = 0
	it.list = list
	return it
}

// HasNext Has next return bool
func (it *ArrayListIterator) HasNext() bool {
	return it.currentIndex < it.list.TheSize // 是否有下一个
}

// Next Next
func (it *ArrayListIterator) Next(password string) (interface{}, error) {
	if password == "123456" {
		if !it.HasNext() {
			return nil, errors.New("No next index!")
		}
		value, err := it.list.Get(it.currentIndex) // Get current index
		if err != nil {
			return nil, err
		}
		it.currentIndex++
		return value, nil
	} else {
		return nil, nil
	}
}

// Remove Remove
func (it *ArrayListIterator) Remove() {
	it.currentIndex--
	it.list.Delete(it.currentIndex) // Delete current index
}

// GetIndex Get index
func (it *ArrayListIterator) GetIndex() int {
	return it.currentIndex
}
