package Variant

type optionTag uint

const (
	SomeTag optionTag = iota
	NoneTag
)

type Option[T any] interface {
	tag() optionTag
}

type Some[T any] struct {
	Value T
}

type None struct {
}

func NewSome[T any](a T) Option[T] {
	return Some[T]{Value: a}
}

func NewNone[T any]() Option[T] { return None{} }

func (s Some[T]) tag() optionTag { return SomeTag }

func (n None) tag() optionTag {
	return NoneTag
}

type eitherTag int

const (
	LeftTag eitherTag = iota
	RightTag
)

type Either interface {
	tag() eitherTag
}

type Left[T any] struct {
	Value T
}

type Right[T any] struct {
	Value T
}

func (l Left[T]) tag() eitherTag {
	return LeftTag
}

func (r Right[T]) tag() eitherTag { return RightTag }
