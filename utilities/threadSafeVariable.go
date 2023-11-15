package utilities

import (
	"sync"

	"golang.org/x/exp/constraints"
)

type ThreadVariable[T any] struct {
	value T
	mu    sync.Mutex
}

func (t *ThreadVariable[T]) Get() T {
	t.mu.Lock()
	ret := t.value
	t.mu.Unlock()
	return ret
}

func (t *ThreadVariable[T]) Set(val T) {
	t.mu.Lock()
	t.value = val
	t.mu.Unlock()
}

type ThreadNumericVariable[T constraints.Integer | constraints.Float | constraints.Complex] struct {
	value T
	mu    sync.Mutex
}

func (t *ThreadNumericVariable[T]) Get() T {
	t.mu.Lock()
	ret := t.value
	t.mu.Unlock()
	return ret
}

func (t *ThreadNumericVariable[T]) Set(val T) {
	t.mu.Lock()
	t.value = val
	t.mu.Unlock()
}

func (t *ThreadNumericVariable[T]) Add1(val T) {
	t.mu.Lock()
	t.value++
	t.mu.Unlock()
}
