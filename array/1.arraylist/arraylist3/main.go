package ArrayList

import (
	"errors"
	"fmt"
)

// 接口定义
type List interface {
	Size() int                                  // 数组大小
	Get(index int) (interface{}, error)         // 获取元素
	Set(index int, newval interface{}) error    // 提交数据
	Insert(index int, newval interface{}) error // 修改数据
	Append(newval interface{})                  // 追加
	Clear()                                     //清空
	Delete(index int) error                     // 删除
	String() string                             // 返回字符串
}

// 数据结构定义
type ArrayList struct {
	dataStore []interface{} // 数组存储
	theSize   int           // 数组大小
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)                      // 初始化结构体
	list.dataStore = make([]interface{}, 0, 10) // 开辟空间10个
	list.theSize = 10
	return list
}

func (list *ArrayList) Size() int {
	return list.theSize
}

//
func (list *ArrayList) Append(newval interface{}) {
	list.dataStore = append(list.dataStore, newval) // 叠加数据
	list.theSize++
}

// 抓去数据
func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.theSize {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

// 返回字符串
func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}
