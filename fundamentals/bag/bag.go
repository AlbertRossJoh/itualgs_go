package bag

import (
	. "github.com/AlbertRossJoh/itualgs_go/utilities"
)

type Bag[T any] struct {
	Items []T
	Size  int
}

func New[T any]() *Bag[T] {
	return &Bag[T]{
		Items: make([]T, 0, 16),
		Size:  0,
	}
}

func (b *Bag[T]) Add(item T) {
	b.Items = append(b.Items, item)
	b.Size++
}

func (b *Bag[T]) IsEmpty() bool {
	return b.Size == 0
}

func (b *Bag[T]) GetIterator() *Iterator[T] {
	return NewIterator(&b.Items)
}
