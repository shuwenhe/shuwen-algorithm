package arraylist

type StackArray interface {
	Clear()                // 清空
	Size() int             // 大小
	Pop() interface{}      // 弹出
	Push(data interface{}) // 压入
	IsFull() bool          // 判断是否已满
	IsEmpty() bool         // 判断是否已空
}

// Stack Struct
type Stack struct {
	myArray *ArrayList
	capsize int // Caption
}

// NewStack New Stack
func NewArrayListStack() *Stack {
	myStack := new(Stack)
	myStack.myArray = NewArrayList()
	myStack.capsize = 10 // Caption
	return myStack
}

// Clear 清空
func (myStack *Stack) Clear() {
	myStack.myArray.Clear()
	myStack.capsize = 10 // Cation size
}

// Size Stack size
func (myStack *Stack) Size() int {
	return myStack.myArray.TheSize
}

// Pop Pop
func (myStack *Stack) Pop() interface{} {
	if !myStack.IsEmpty() {
		last := myStack.myArray.dataStore[myStack.myArray.TheSize-1] // Last data
		myStack.myArray.Delete(myStack.myArray.TheSize - 1)          // Delete last data
		return last
	}
	return nil
}

// Push Push
func (myStack *Stack) Push(data interface{}) {
	if !myStack.IsFull() {
		myStack.myArray.Append(data)
		myStack.myArray.TheSize++
	}
}

// IsFull 判断是否已满
func (myStack *Stack) IsFull() bool {
	if myStack.myArray.TheSize >= myStack.capsize {
		return true
	}
	return false
}

// IsEmpty 判断是否已空
func (myStack *Stack) IsEmpty() bool {
	if myStack.myArray.TheSize == 0 {
		return true
	}
	return false
}
