package main

import (
	"fmt"
)

func max[T comparable](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func main() {
	fmt.Println("Max of 3 and 5:", max(3, 5))
	fmt.Println("Max of 'a' and 'b':", max('a', 'b'))

	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	fmt.Println("Popped from int stack:", intStack.Pop())
	fmt.Println("Popped from int stack:", intStack.Pop())

	stringStack := Stack[string]{}
	stringStack.Push("Hello")
	stringStack.Push("World")

	fmt.Println("Popped from string stack:", stringStack.Pop())
	fmt.Println("Popped from string stack:", stringStack.Pop())
}
