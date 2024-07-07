package stack

type IStack interface {
	New()
	Push(value interface{})
	Pop() interface{}
	Top()	interface{}
}

type Stack struct {
	Point int
}

func (s *Stack) Push() {
	s.Point++
}

func (s *Stack) Pop() interface{} {
	s.Point--
	return 0
}

func (s *Stack) Top() interface{} {
	return 0
}
