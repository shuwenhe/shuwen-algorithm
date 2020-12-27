package main

type MyStack struct {
	queue []int
}

func InitStack() (s MyStack) {
	return
}

func (s *MyStack) Push(x int) {
	n := len(s.queue)
	s.queue = append(s.queue, x)
	for ; n > 0; n-- {
		s.queue = append(s.queue, s.queue[0])
		s.queue = s.queue[1:]
	}
}

func (s *MyStack) Pop() int {
	v := s.queue[0]
	s.queue = s.queue[1:]
	return v
}

func (s *MyStack) Top() int {
	return s.queue[0]
}

func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}

func main() {

}
