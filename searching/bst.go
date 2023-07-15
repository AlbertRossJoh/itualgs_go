package searching

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type node[K constraints.Ordered, T any] struct {
	Value T
	Key   K
	Left  *node[K, T]
	Right *node[K, T]
}

type BST[K constraints.Ordered, T any] struct {
	root *node[K, T]
}

func NewBST[K constraints.Ordered, T any](key K, val T) BST[K, T] {
	return BST[K, T]{&node[K, T]{
		Key:   key,
		Value: val,
		Left:  nil,
		Right: nil,
	}}
}

func (self BST[K, T]) Put(key K, val T) {
	node := &node[K, T]{
		Value: val,
		Key:   key,
		Left:  nil,
		Right: nil,
	}

	if self.root == nil {
		self.root = node
		return
	}

	put[K, T](self.root, node)
}

func put[K constraints.Ordered, T any](parent *node[K, T], node *node[K, T]) {
	if parent == nil {
		parent = node
		return
	}

	if node.Key < parent.Key {
		if parent.Left == nil {
			parent.Left = node
			return
		}
		put[K, T](parent.Left, node)
	} else if node.Key > parent.Key {
		if parent.Right == nil {
			parent.Right = node
			return
		}
		put[K, T](parent.Right, node)
	} else {
		parent.Key = node.Key
		parent.Value = node.Value
		return
	}
}

func (self BST[K, T]) Get(key K) (T, error) {
	if self.root == nil {
		var res T
		return res, errors.New("The tree is ")
	}

	isCorrect, node := walk[K, T](key, self.root)

	if isCorrect {
		return node.Value, nil
	}

	var res T
	return res, errors.New("The key could not be found")

}

func walk[K constraints.Ordered, T any](key K, node *node[K, T]) (bool, *node[K, T]) {
	if node.Key > key {
		if node.Left == nil {
			return false, node
		}
		return walk(key, node.Left)
	} else if node.Key < key {
		if node.Right == nil {
			return false, node
		}
		return walk(key, node.Right)
	} else {
		return true, node
	}
}

func (self BST[K, T]) Delete(key K) (T, error) {

	isCorrect, node := walk[K, T](key, self.root)

	if !isCorrect {
		var res T
		return res, errors.New("The key could not be found")
	}
	ret := node.Value
	*self.root = *self._delete(key, self.root)
	return ret, nil
}

func (self BST[K, T]) _delete(key K, node *node[K, T]) *node[K, T] {
	if node == nil {
		return nil
	}

	var cmp = node.cmpKey(key)
	if cmp < 0 {
		node.Left = self._delete(key, node.Left)
	} else if cmp > 0 {
		node.Right = self._delete(key, node.Right)
	} else {
		if node.Right == nil {
			return node.Left
		}
		if node.Left == nil {
			return node.Right
		}

		tmp := *node
		node = tmp.Right.Min()
		node.Right = deleteMin(tmp.Right)
		node.Left = tmp.Left
	}
	return node
}

func (self node[K, T]) cmpKey(key K) int8 {
	if self.Key < key {
		return 1
	} else if self.Key > key {
		return -1
	}
	return 0
}

func (self BST[K, T]) Min() K {
	return self.root.Min().Key
}

func (self node[K, T]) Min() *node[K, T] {
	curr := &self
	next := curr

	for next != nil {
		curr = next
		next = next.Left
	}
	return curr
}

func (self BST[K, T]) Max() K {
	return self.root.Max().Key
}

func (self node[K, T]) Max() *node[K, T] {
	curr := &self
	next := curr

	for next != nil {
		curr = next
		next = next.Right
	}
	return curr
}

func (self BST[K, T]) DeleteMax() {
	self.root = deleteMax(self.root)
}

func deleteMax[K constraints.Ordered, T any](node *node[K, T]) *node[K, T] {
	if node.Right == nil {
		return node.Left
	}
	node.Right = deleteMin(node.Right)
	return node
}

func (self BST[K, T]) DeleteMin() {
	self.root = deleteMin(self.root)
}

func deleteMin[K constraints.Ordered, T any](node *node[K, T]) *node[K, T] {
	if node.Left == nil {
		return node.Right
	}
	node.Left = deleteMin(node.Left)
	return node
}

func (self BST[K, T]) Floor(key K) (K, error) {
	if self.root == nil {
		var ret K
		return ret, errors.New("The tree is empty")
	}

	ret := floor[K, T](key, self.root)
	if ret == nil {
		var ret K
		return ret, errors.New("The key could not be found")
	}
	return ret.Key, nil
}

func floor[K constraints.Ordered, T any](key K, node *node[K, T]) *node[K, T] {
	if node == nil {
		return nil
	}
	cmp := node.cmpKey(key) * -1

	if cmp == 0 {
		return node
	}

	if cmp < 0 {
		return floor(key, node.Left)
	}

	tmp := floor(key, node.Right)

	if tmp != nil {
		return tmp
	} else {
		return node
	}
}

func (self BST[K, T]) Ceiling(key K) (K, error) {
	if self.root == nil {
		var ret K
		return ret, errors.New("The tree is empty")
	}
	ret := ceiling[K, T](key, self.root)
	if ret == nil {
		var ret K
		return ret, errors.New("The key could not be found")
	}
	return ret.Key, nil
}

func ceiling[K constraints.Ordered, T any](key K, node *node[K, T]) *node[K, T] {
	if node == nil {
		return nil
	}
	cmp := node.cmpKey(key) * -1

	if cmp == 0 {
		return node
	}

	if cmp < 0 {
		tmp := ceiling(key, node.Left)
		if tmp != nil {
			return tmp
		} else {
			return node
		}
	}

	return ceiling(key, node.Right)
}
