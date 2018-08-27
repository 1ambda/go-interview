package codility

// https://app.codility.com/programmers/lessons/7-stacks_and_queues/nesting/
// expected worst-case time complexity is O(N);
// expected worst-case space complexity is O(1) (not counting the storage required for input arguments).

type StringStack struct {
	values []string
	top int // index of stack's top
}

func NewStringStack() *StringStack {
	return &StringStack{
		values: make([]string, 0),
		top: -1,
	}
}

// push and return the length of this stack
func (s *StringStack) Push(value string) int {
	s.values = append(s.values, value)
	s.top += 1

	return len(value)
}

func (s *StringStack) Pop() string {
	if len(s.values) == 0 {
		panic("Can't execute Pop operation on an empty stack")
	}

	value := s.values[len(s.values) - 1]
	s.values = s.values[:len(s.values) - 1]

	return value
}

func (s *StringStack) IsEmpty() bool {
	return len(s.values) == 0
}



func Solution(S string) int {
	stack := NewStringStack()

	for _, rune := range S {
		char := string(rune)

		if char == "(" {
			stack.Push(char)

		} else {
			if stack.IsEmpty() {
				return 0
			}

			stack.Pop()
		}
	}

	if stack.IsEmpty() {
		return 1
	}

	return 0
}

func PoorSolution(S string) int {

	stack := make([]string, 0)

	for _, rune := range S {
		char := string(rune)

		if char == "(" {
			// push: bad performance
			// stack = append([]string{char}, stack...)

			// push: better, not still not O(1)
			stack = append(stack, "")
			copy(stack[1:], stack)
			stack[0] = char
		} else {
			if len(stack) == 0 {
				return 0
			}

			// pop
			stack = stack[1:]
		}
	}

	if len(stack) == 0 {
		return 1
	}

	return 0
}


