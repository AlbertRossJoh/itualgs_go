package stack

import (
	"github.com/AlbertRossJoh/itualgs_go/customerrors"
	util "github.com/AlbertRossJoh/itualgs_go/utilities"
)

type Stack[T any] struct {
	items []T
}

func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0, size),
	}
}

func NewEmptyStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0, 16),
	}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Size() == 0 {
		return *new(T), &customerrors.ErrEmptyStack{}
	}
	ret := s.items[s.Size()-1]
	s.items = s.items[:s.Size()-1]
	return ret, nil
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Peek() (T, error) {
	if s.Size() == 0 {
		return *new(T), &customerrors.ErrEmptyStack{}
	}
	return s.items[s.Size()-1], nil
}

func (s *Stack[T]) Clear() {
	s.items = make([]T, 0, s.Size())
}

func (s *Stack[T]) Items() []T {
	return s.items
}

func (s *Stack[T]) GetIterator() *util.Iterator[T] {
	tmp := make([]T, s.Size())
	copy(tmp, s.items)
	return util.NewIterator[T](&tmp)
}
