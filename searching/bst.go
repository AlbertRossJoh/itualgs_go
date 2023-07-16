package searching

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type node[K constraints.Ordered, T any] struct {
	value T
	key   K
	left  *node[K, T]
	right *node[K, T]
}

type BST[K constraints.Ordered, T any] struct {
	root *node[K, T]
}

func NewBST[K constraints.Ordered, T any](key K, val T) BST[K, T] {
	return BST[K, T]{&node[K, T]{
		key:   key,
		value: val,
		left:  nil,
		right: nil,
	}}
}

func NewEmptyBST[K constraints.Ordered, T any]() BST[K, T] {
	return BST[K, T]{nil}
}

func (self *BST[K, T]) Put(key K, val T) {
	node := &node[K, T]{
		value: val,
		key:   key,
		left:  nil,
		right: nil,
	}

	if self.root == nil {
		*self.root = *node
		return
	}

	put[K, T](self.root, node)
}

func put[K constraints.Ordered, T any](parent *node[K, T], node *node[K, T]) {
	if parent == nil {
		*parent = *node
		return
	}

	if node.key < parent.key {
		if parent.left == nil {
			parent.left = node
			return
		}
		put[K, T](parent.left, node)
	} else if node.key > parent.key {
		if parent.right == nil {
			parent.right = node
			return
		}
		put[K, T](parent.right, node)
	} else {
		parent.key = node.key
		parent.value = node.value
		return
	}
}

func (self *BST[K, T]) Get(key K) (T, error) {
	if self.root == nil {
		var res T
		return res, errors.New("the tree is empty")
	}

	isCorrect, node := walk[K, T](key, self.root)

	if isCorrect {
		return node.value, nil
	}

	var res T
	return res, errors.New("the key could not be found")

}

func walk[K constraints.Ordered, T any](key K, node *node[K, T]) (bool, *node[K, T]) {
	if node.key > key {
		if node.left == nil {
			return false, node
		}
		return walk(key, node.left)
	} else if node.key < key {
		if node.right == nil {
			return false, node
		}
		return walk(key, node.right)
	} else {
		return true, node
	}
}

func (self *BST[K, T]) Delete(key K) (T, error) {
	if self.root == nil {
		var res T
		return res, errors.New("the tree is empty")
	}

	isCorrect, node := walk[K, T](key, self.root)

	if !isCorrect {
		var res T
		return res, errors.New("the key could not be found")
	}
	ret := node.value
	self.root = self._delete(key, self.root)
	return ret, nil
}

func (self *BST[K, T]) _delete(key K, node *node[K, T]) *node[K, T] {
	if node == nil {
		return nil
	}

	var cmp = node.cmpkey(key)
	if cmp < 0 {
		node.left = self._delete(key, node.left)
	} else if cmp > 0 {
		node.right = self._delete(key, node.right)
	} else {
		if node.right == nil {
			return node.left
		}
		if node.left == nil {
			return node.right
		}

		tmp := *node
		node = tmp.right.Min()
		node.right = deleteMin(tmp.right)
		node.left = tmp.left
	}
	return node
}

func (self *node[K, T]) cmpkey(key K) int8 {
	if self.key < key {
		return 1
	} else if self.key > key {
		return -1
	}
	return 0
}

func (self *BST[K, T]) Min() (K, error) {
	if self.root == nil {
		var ret K
		return ret, errors.New("the tree is empty")
	}
	ret := self.root.Min()
	return ret.key, nil
}

func (self node[K, T]) Min() *node[K, T] {
	curr := &self
	next := curr

	for next != nil {
		curr = next
		next = next.left
	}
	return curr
}

func (self *BST[K, T]) Max() (K, error) {
	if self.root == nil {
		var ret K
		return ret, errors.New("the tree is empty")
	}
	ret := self.root.Max()
	return ret.key, nil
}

func (self node[K, T]) Max() *node[K, T] {
	curr := &self
	next := curr

	for next != nil {
		curr = next
		next = next.right
	}
	return curr
}

func (self *BST[K, T]) DeleteMax() {
	if self.root == nil {
		panic("the tree is empty")
	}
	self.root = deleteMax(self.root)
}

func deleteMax[K constraints.Ordered, T any](node *node[K, T]) *node[K, T] {
	if node.right == nil {
		return node.left
	}
	node.right = deleteMax(node.right)
	return node
}

func (self *BST[K, T]) DeleteMin() {
	if self.root == nil {
		panic("the tree is empty")
	}
	self.root = deleteMin(self.root)
}

func deleteMin[K constraints.Ordered, T any](node *node[K, T]) *node[K, T] {
	if node.left == nil {
		return node.right
	}
	node.left = deleteMin(node.left)
	return node
}

func (self *BST[K, T]) Floor(key K) (K, error) {
	if self.root == nil {
		var ret K
		return ret, errors.New("the tree is empty")
	}

	ret := floor[K, T](key, self.root)
	if ret == nil {
		var ret K
		return ret, errors.New("the key could not be found")
	}
	return ret.key, nil
}

func floor[K constraints.Ordered, T any](key K, node *node[K, T]) *node[K, T] {
	if node == nil {
		return nil
	}
	cmp := node.cmpkey(key) * -1

	if cmp == 0 {
		return node
	}

	if cmp < 0 {
		return floor(key, node.left)
	}

	tmp := floor(key, node.right)

	if tmp != nil {
		return tmp
	} else {
		return node
	}
}

func (self *BST[K, T]) Ceiling(key K) (K, error) {
	if self.root == nil {
		var ret K
		return ret, errors.New("the tree is empty")
	}
	ret := ceiling[K, T](key, self.root)
	if ret == nil {
		var ret K
		return ret, errors.New("the key could not be found")
	}
	return ret.key, nil
}

func ceiling[K constraints.Ordered, T any](key K, node *node[K, T]) *node[K, T] {
	if node == nil {
		return nil
	}
	cmp := node.cmpkey(key) * -1

	if cmp == 0 {
		return node
	}

	if cmp < 0 {
		tmp := ceiling(key, node.left)
		if tmp != nil {
			return tmp
		} else {
			return node
		}
	}

	return ceiling(key, node.right)
}

func (self *BST[K, T]) GetAllKeys() ([]K, error) {
	if self.root == nil {
		return nil, errors.New("the tree is empty")
	}
	min, err := self.Min()
	max, err2 := self.Max()
	if err != nil {
		return nil, err
	}
	if err2 != nil {
		return nil, err2
	}
	ret, _ := self.GetkeysInRange(min, max)
	return ret, nil
}

func (self *BST[K, T]) GetkeysInRange(lo K, hi K) ([]K, error) {
	if self.root == nil {
		return nil, errors.New("the tree is empty")
	}
	if lo > hi {
		return nil, errors.New("the low key is greater than the high key")
	}
	var arr []K
	var tmp []T
	keys(self.root, &arr, &tmp, lo, hi)
	return arr, nil
}

func keys[K constraints.Ordered, T any](node *node[K, T], listkeys *[]K, listvalues *[]T, lo K, hi K) {
	if node == nil {
		return
	}
	if lo < node.key {
		keys(node.left, listkeys, listvalues, lo, hi)
	}
	if lo <= node.key && hi >= node.key {
		*listkeys = append(*listkeys, node.key)
		*listvalues = append(*listvalues, node.value)
	}
	if hi > node.key {
		keys(node.right, listkeys, listvalues, lo, hi)
	}
}

func (self *BST[K, T]) IsEmpty() bool {
	return self.root == nil
}

func (self *BST[K, T]) GetAllValues() ([]T, error) {
	if self.root == nil {
		return nil, errors.New("the tree is empty")
	}
	min, err := self.Min()
	max, err2 := self.Max()
	if err != nil {
		return nil, err
	}
	if err2 != nil {
		return nil, err2
	}
	var list []T
	var tmp []K
	keys(self.root, &tmp, &list, min, max)
	return list, nil
}

func (self *BST[K, T]) Mapify() (map[K]T, error) {
	if self.root == nil {
		return nil, errors.New("the tree is empty")
	}

	min, err := self.Min()
	max, err2 := self.Max()
	if err != nil {
		return nil, err
	}
	if err2 != nil {
		return nil, err2
	}

	var allkeys []K
	var allvalues []T
	keys(self.root, &allkeys, &allvalues, min, max)
	ret := make(map[K]T)

	for i := 0; i < len(allkeys); i++ {
		ret[allkeys[i]] = allvalues[i]
	}

	return ret, nil
}
