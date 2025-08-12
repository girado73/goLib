// Package option is a library to extend the functionality of the Go programming language
// It orients itself on the haskell base library and provides similar functions
package option

type Option[T any] struct {
	value T
	ok    bool
}

func Some[T any](v T) Option[T] {
	return Option[T]{value: v, ok: true}
}

func None[T any]() Option[T] {
	var zero T
	return Option[T]{value: zero, ok: false}
}

func (o Option[T]) IsSome() bool {
	return o.ok
}

func (o Option[T]) IsNone() bool {
	return !o.ok
}

func (o Option[T]) Unwrap() T {
	if !o.ok {
		panic("called Unwrap on None value")
	}
	return o.value
}

func (o Option[T]) UnwrapOr(defaultValue T) T {
	if o.ok {
		return o.value
	}
	return defaultValue
}
