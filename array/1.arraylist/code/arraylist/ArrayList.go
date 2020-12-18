package arraylist

import (
	"errors"
	"fmt"
)

// List 接口
type List interface {
	Size() int                                  // 数组大小
	Get(index int) (interface{}, error)         // 抓去第几个元素
	Set(index int, newval interface{}) error    // 修改数据
	Insert(index int, newval interface{}) error // 插入数据
	Append(newval interface{})                  // 追加
	Clear()                                     // 清空
	Delete(index int) error                     // 删除
	String() string                             // 返回字符串
	Iterator() Iterator                         // 接口
}

// ArrayList 数据结构，字符串，整数，实数，泛型
type ArrayList struct {
	dataStore []interface{} // 数组存储
	TheSize   int           // 数组的大小
}

// NewArrayList 开辟内存
func NewArrayList() *ArrayList {
	list := new(ArrayList)                      // 初始化结构体
	list.dataStore = make([]interface{}, 0, 10) // 开辟空间10个
	list.TheSize = 0
	return list
}

// 判断空间是否已经满了
func (list *ArrayList) checkisFull() {
	if list.TheSize == cap(list.dataStore) { // 判断内存使用
		fmt.Println("A", list.dataStore)
		// make 中间的参数，0，没有开辟内存
		newdataStore := make([]interface{}, 2*list.TheSize, 2*list.TheSize) // 开辟双倍内存
		copy(newdataStore, list.dataStore)                                  // 拷贝:把list.dateStore copy to newdataStore
		// for i := 0; i < len(list.dataStore); i++ {
		// 	newdataStore[i] = list.dataStore[i]
		// }
		list.dataStore = newdataStore // 赋值
		fmt.Println("A", list.dataStore)
		fmt.Println("A", newdataStore)
	}
}

// Size 数组大小
func (list *ArrayList) Size() int {
	return list.TheSize // 返回数据大小
}

// Append 追加
func (list *ArrayList) Append(newval interface{}) { // 追加
	list.dataStore = append(list.dataStore, newval) // 叠加数据
	list.TheSize++
}

// Get 抓去数据
func (list *ArrayList) Get(index int) (interface{}, error) { // 抓去第几个元素
	if index < 0 || index >= list.TheSize {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

// ArrayList 返回数据字符串
func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}

// Set 修改数据
func (list *ArrayList) Set(index int, newval interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.dataStore[index] = newval // 设置
	return nil
}

// Insert 插入数据
func (list *ArrayList) Insert(index int, newval interface{}) error { // 插入数据
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.checkisFull()                               // 监测内存，如果满了，自动追加
	list.dataStore = list.dataStore[:list.TheSize+1] // 插入数据，内存移动一位
	for i := list.TheSize; i > index; i-- {          // 从后往前移动
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newval // 插入数据
	list.TheSize++                 // 索引追加

	return nil
}

// Clear 清空
func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10) // 开辟空间10个
	list.TheSize = 0
}

// Delete 删除
func (list *ArrayList) Delete(index int) error {
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...) // 重新叠加，跳过中间一个
	list.TheSize--
	return nil
}
