// Package queue
// This is a channel based queue implementation.
// This queue is strictly to be use when needing a thread safe queue.
// As the queue is not resizing the initial memory allocation will be persistent until the queue is closed.
package queue

import (
	"github.com/AlbertRossJoh/itualgs_go/customerrors"
)

type ChannelQueue[T any] struct {
	items chan T
	size  int
}

func NewBufQueue[T any](size int) *ChannelQueue[T] {
	return &ChannelQueue[T]{
		items: make(chan T, size),
		size:  0,
	}
}

func (q *ChannelQueue[T]) Enqueue(item T) error {
	if q.size == cap(q.items) {
		return &customerrors.ErrQueueFull{}
	}
	q.items <- item
	q.size++
	return nil
}

func (q *ChannelQueue[T]) Dequeue() (T, error) {
	if q.size == 0 {
		return *new(T), &customerrors.ErrQueueEmpty{}
	}
	q.size--
	return <-q.items, nil
}

func (q *ChannelQueue[T]) Size() int {
	return q.size
}

func (q *ChannelQueue[T]) Cap() int {
	return cap(q.items)
}

func (q *ChannelQueue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *ChannelQueue[T]) IsFull() bool {
	return q.size == cap(q.items)
}

func (q *ChannelQueue[T]) Close() {
	close(q.items)
}

func (q *ChannelQueue[T]) Items() <-chan T {
	return q.items
}
