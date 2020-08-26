package stackarray

type StackArray interface {
	Clear()                // 清空
	Size() int             // 大小
	Pop() interface{}      // 弹出
	Push(data interface{}) // 压入
	IsFull() bool          // 是否满了
	IsEmpty() bool         // 是否为空
}

type Stack struct {
	dataSource  []interface{}
	capSize     int // 最大范围
	currentSize int // 当前实际使用大小
}

func NewStack() *Stack {
	mystack := new(Stack)
	mystack.dataSource = make([]interface{}, 0, 10)
}

func (s *Stack) Clear() { // 清空

}

func (s *Stack) Size() int { // 大小

}

func (s *Stack) Pop() interface{} { // 弹出

}

func (s *Stack) Push(data interface{}) { // 压入
}

func (s *Stack) IsFull() bool { // 是否满了
}

func (s *Stack) IsEmpty() bool { // 是否为空
}
