package main

// create a stack struct
type Stack[T any] struct {
	elements []T
}

// push add elements in stack
func (s *Stack[T]) Push(item T) {
	s.elements = append(s.elements, item)
}

// pop element in stack
func (s *Stack[T]) Pop() T {
	length := len(s.elements)
	if length == 0 {
		panic("stack is empty")
	}
	popItem := s.elements[length-1]
	s.elements = s.elements[:length-1]
	return popItem
}
