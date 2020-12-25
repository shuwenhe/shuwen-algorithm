package linkstack

// Node struct
type Node struct { // 1->2->3->...->n
	Data  interface{}
	pNext *Node // Shuwen->Richard ; Shuwen->Ritchie
}

// LinkStack interface
type LinkStack interface {
	IsEmpty() bool         // Whether is empty (IsFull)
	Push(data interface{}) // Push data into stack
	Pop() interface{}      // Pop data out stack
	Length() int           // The length of the stack
}
