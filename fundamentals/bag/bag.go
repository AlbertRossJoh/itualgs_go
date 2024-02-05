package bag

import (
	util "github.com/AlbertRossJoh/itualgs_go/utilities/sharedFunctions"
)

type Bag[T any] struct {
	Items []T
	Size  int
}

func NewBag[T any]() *Bag[T] {
	return &Bag[T]{
		Items: make([]T, 0, 16),
		Size:  0,
	}
}

func (b *Bag[T]) Add(item ...T) {
	b.Items = append(b.Items, item...)
	b.Size++
}

func (b *Bag[T]) IsEmpty() bool {
	return b.Size == 0
}

func (b *Bag[T]) GetIterator() *util.Iterator[T] {
	return util.NewIterator(&b.Items)
}

func (b *Bag[T]) Clone() Bag[T] {
	tmp := make([]T, b.Size)
	copy(tmp, b.Items)
	return Bag[T]{
		Items: tmp,
		Size:  b.Size,
	}
}
