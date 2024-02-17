package stack

func New[T any]() Stack[T] {
	s := make([]T, 0, 16)
	return Stack[T]{a: s, t: -1, Len: 0}
}

type Stack[T any] struct {
	a   []T
	t   int
	Len int
}

func (s *Stack[T]) Push(item T) {
	s.t++
	s.Len++

	c := cap(s.a)
	n := c

	if s.t >= c {
		c *= 2
		a := s.a
		s.a = make([]T, n, c)
		copy(a, s.a)
	}
	s.a[s.t] = item
}

func (s *Stack[T]) Empty() bool {
	return s.t < 0
}

func (s *Stack[T]) Pop() Option[T] {
	if s.t < 0 {
		return None[T]{}
	}

	item := s.a[s.t]
	s.t--

	return Some[T]{item: item}
}

func (s *Stack[T]) Peek() Option[T] {
	if s.t < 0 {
		return None[T]{}
	}

	return Some[T]{item: s.a[s.t]}
}
