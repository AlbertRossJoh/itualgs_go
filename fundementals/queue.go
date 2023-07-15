package fundementals

type Queue[T any] struct {
	items []T
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0, size),
	}
}

func NewEmptyQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0, 16),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() T {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *Queue[T]) Peek() T {
	return q.items[0]
}

func (q *Queue[T]) Clear() {
	q.items = make([]T, q.Size())
}

func (q *Queue[T]) Items() []T {
	return q.items
}
