package main

import (
	"errors"
	"fmt"
)

func main() {
	// list := NewArrayList()
	var list List = NewArrayList() // 定义接口对象，赋值的对象必须实现接口的所有方法
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append("a")
	list.Append("b")
	list.Append("c")
	fmt.Println("list = ", list)
}

// List 接口
type List interface {
	Size() int                          // 数组大小
	Get(index int) (interface{}, error) // 获取第几个元素
	Append(newval interface{})          // 追加数据
	String() string                     // 返回字符串

	Set(index int, newval interface{}) error    // 修改数据
	Insert(index int, newval interface{}) error // 插入数据
	Clear()                                     // 清空数据
	Delete(index int) error                     // 删除第几个数据
}

// ArrayList struct
type ArrayList struct {
	dataStore []interface{}
	TheSize   int // 数组的大小
}

// NewArrayList New array list
func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.dataStore = make([]interface{}, 0, 10) // 开辟10个空间
	list.TheSize = 0                            // 初始化为0
	return list
}

// Size ArrayList size
func (list *ArrayList) Size() int {
	return list.Size() // Return data size
}

// Get 获取第几个元素
func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index > list.Size() {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

// Append 追加数据
func (list *ArrayList) Append(newval interface{}) {
	list.dataStore = append(list.dataStore, newval) // Add data
	list.TheSize++
}

// String 返回字符串
func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}

// Set 修改数据
func (list *ArrayList) Set(index int, newval interface{}) error {
	if index < 0 || index > list.TheSize {
		return errors.New("索引越界")
	}
	list.dataStore[index] = newval // Set
	return nil
}

// Insert Insert data
func (list *ArrayList) Insert(index int, newval interface{}) error {
	if index < 0 || index > list.TheSize {
		return errors.New("索引越界")
	}
	return nil
}

// Clear Clear data
func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)
	list.TheSize = 0
}

// Delete Delete第几个data
func (list *ArrayList) Delete(index int) error {
	return nil
}
