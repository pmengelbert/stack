package stack

type Option[T any] interface {
	Unwrap() T
	IsSome() bool
}

type None[T any] struct{}

func (_ None[T]) Unwrap() T {
	panic("called `unwrap` on a None value")
}

func (_ None[T]) IsSome() bool {
	return false
}

func newNone[T any]() Option[T] {
	return None[T]{}
}

type Some[T any] struct {
	item T
}

func (s Some[T]) Unwrap() T {
	return s.item
}

func (_ Some[T]) IsSome() bool {
	return true
}

func newSome[T any](item T) Option[T] {
	return Some[T]{item: item}
}
