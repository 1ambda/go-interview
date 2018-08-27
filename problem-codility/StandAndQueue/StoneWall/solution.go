package codility

type IntStack struct {
	values []int
	top    int // index of stack's top
}

func NewIntStack() *IntStack {
	return &IntStack{
		values: make([]int,0),
		top: -1,
	}
}

func (s *IntStack) Top() int {
	return s.values[s.top]
}

func (s *IntStack) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *IntStack) Push(value int) int {
	s.values = append(s.values, value)
	s.top += 1

	return len(s.values)
}

func (s *IntStack) Pop() int {
	if s.IsEmpty() {
		panic("Can't pop on an empty stack")
	}

	value := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	s.top -= 1

	return value
}

func Solution(H []int) int {

	stack := NewIntStack()
	count := 0

	for _, h := range H {
		// remove higher blocks since they are already counted
		for !stack.IsEmpty() && stack.Top() > h {
			stack.Pop()
		}

		if stack.IsEmpty() || stack.Top() < h {
			stack.Push(h)
			count += 1
		}
	}

	return count
}
