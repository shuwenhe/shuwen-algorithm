package arraylist

// Iterator Iterator interface
type Iterator interface {
	HasNext() bool              // 是否有下一个
	Next() (interface{}, error) // 下一个
	Remove()                    // Delete
	GetIndex()                  // Get index
}

// Iterable Iterable interface
type Iterable interface {
	Iterator() Iterator // 构造初始化接口
}

// ArrayIterator Array Iterator struct 构造指针访问数组
type ArrayIterator struct {
	list          *ArrayList // 数组指针
	currentIndext int        // Current Index
}

// Iterator Iterator
func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayIterator) //
	it.currentIndext = 0
	it.list = list
	return it
}

// HasNext 是否有下一个
func (it *ArrayIterator) HasNext() bool {
	return it.currentIndext < it.list.TheSize // 是否有下一个
}

// Next 下一个
func (it *ArrayIterator) Next() (interface{}, error) {
	return nil, nil
}

// Remove Remove
func (it *ArrayIterator) Remove() {

}

// GetIndex Get index
func (it *ArrayIterator) GetIndex() {

}
