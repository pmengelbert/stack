package stack

import "fmt"

type Stack[T any] struct {
	a []T
}

func (s *Stack[T]) Push(item T) {
	s.a = append(s.a, item)
}

func (s *Stack[T]) Empty() bool {
	return len(s.a) == 0
}

func (s *Stack[T]) Len() int {
	return len(s.a)
}

func (s *Stack[T]) Pop() T {
	if len(s.a) == 0 {
		panic(fmt.Sprintf("stack is empty: %#v", *s))
	}

	item := s.a[len(s.a)-1]
	s.a = s.a[:len(s.a)-1]

	return item
}
