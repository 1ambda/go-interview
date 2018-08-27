package codility

type IntStack struct {
	values []int
	top    int // index of stack's top
}

func NewIntStack() *IntStack {
	return &IntStack{
		values: make([]int, 0),
		top:    -1,
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

func (s *IntStack) Count() int {
	return len(s.values)
}

func Solution(A []int, B []int) int {

	downstreamFishSizes := NewIntStack()
	alive := 0

	for i := range A {
		size := A[i]
		direction := B[i]

		if direction == 1 {
			downstreamFishSizes.Push(size)
		} else {
			for !downstreamFishSizes.IsEmpty() && downstreamFishSizes.Top() < size {
				downstreamFishSizes.Pop()
			}

			if downstreamFishSizes.IsEmpty() {
				alive += 1
			}
		}
	}

	return downstreamFishSizes.Count() + alive
}
