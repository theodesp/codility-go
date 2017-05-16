package stacks_queues

type StringStack struct {
	size int
	data []string
}

func NewStringStack(len int) *StringStack {
	return &StringStack{
		size: 0,
		data: make([]string, len),
	}
}

func (s *StringStack) Push(item string) {
	if s.size < len(s.data) {
		s.data[s.size] = item
		s.size++
	}
}

func (s *StringStack) Pop() string {
	if s.size > 0 {
		item := s.data[s.size - 1]
		s.size -= 1

		return item
	} else {
		return ""
	}
}

func (s *StringStack) Front() string {
	if s.size > 0 {
		return s.data[s.size - 1]
	} else {
		return ""
	}
}

type IntStack struct {
	size int
	data []int
}

func NewIntStack(len int) *IntStack {
	return &IntStack{
		size: 0,
		data: make([]int, len),
	}
}

func (s *IntStack) Push(item int) {
	if s.size < len(s.data) {
		s.data[s.size] = item
		s.size++
	}
}

func (s *IntStack) Pop() int {
	item := s.data[s.size - 1]
	s.size -= 1

	return item
}

func (s *IntStack) Front() int {
	return s.data[s.size - 1]
}

