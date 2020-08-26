package arraylist

import (
	"errors"
)

// Iterator interface
type Iterator interface {
	HasNext() bool                             // 是否有下一个
	Next(password string) (interface{}, error) // 下一个
	Remove()                                   // 删除
	GetIndex() int                             // 得到索引
}

// Iterable interface
type Iterable interface {
	Iterator() Iterator // 迭代器设计模式，构造初始化接口
}

// ArrayListIterator struct
type ArrayListIterator struct {
	list         *ArrayList // 数组指针
	currentindex int        // 当前索引
}

func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator) // 构造迭代器
	it.currentindex = 0
	it.list = list
	return it
}

// HasNext 是否有下一个
func (it *ArrayListIterator) HasNext() bool {
	return it.currentindex < it.list.TheSize // 是否有下一个
}

// Next 下一个
func (it *ArrayListIterator) Next(password string) (interface{}, error) {
	if password == "111111" {
		if !it.HasNext() {
			return nil, errors.New("没有下一个")
		}
		value, err := it.list.Get(it.currentindex) // 抓取当前数据
		it.currentindex++
		return value, err
	} else {
		return nil, nil
	}
}

// Remove 删除
func (it *ArrayListIterator) Remove() {
	it.currentindex--
	it.list.Delete(it.currentindex) // 删除一个元素
}

// GetIndex 得到索引
func (it *ArrayListIterator) GetIndex() int {
	return it.currentindex
}
