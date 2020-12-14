package arraylist

import (
	"errors"
	"fmt"
)

// List 接口
type List interface {
	Size() int                                  // 数组大小
	Get(index int) (interface{}, error)         // 获取第几个元素
	Append(newval interface{})                  // 追加数据
	String() string                             // 返回字符串
	Set(index int, newval interface{}) error    // 修改数据
	Insert(index int, newval interface{}) error // 插入数据
	Clear()                                     // 清空数据
	Delete(index int) error                     // 删除第几个数据
}

// ArrayList struct
type ArrayList struct {
	dataStore []interface{} // Array 任意类型的
	TheSize   int           // Array的大小
}

// NewArrayList New array list
func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.dataStore = make([]interface{}, 0, 10) // 开辟10个空间：长度为0，容量为10的slice
	list.TheSize = 0                            // 初始化为0
	return list
}

// CheckIsFull Check the array is full?
func (list *ArrayList) CheckIsFull() {
	if list.TheSize == cap(list.dataStore) { // 判断内存使用
		newdataStore := make([]interface{}, 2*list.TheSize) // 开辟2倍的memory空间
		copy(newdataStore, list.dataStore)                  // Copy
		list.dataStore = newdataStore                       // 赋值
	}
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
	list.CheckIsFull()                               // 检测内存，如果满了自动追加
	list.dataStore = list.dataStore[:list.TheSize+1] // 插入数据，内存需要向后移动一位
	for i := list.TheSize; i > index; i-- {          // 从后面往前面移动
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newval // Insert data
	list.TheSize++                 // 追加索引
	return nil
}

// Clear Clear data
func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)
	list.TheSize = 0
}

// Delete Delete第几个data
func (list *ArrayList) Delete(index int) error {
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...) // 重新叠加跳过index
	list.TheSize--
	return nil
}
