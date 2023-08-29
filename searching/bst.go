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

func (bst *BST[K, T]) Put(key K, val T) {
	node := &node[K, T]{
		value: val,
		key:   key,
		left:  nil,
		right: nil,
	}

	if bst.root == nil {
		*bst.root = *node
		return
	}

	put[K, T](bst.root, node)
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

func (bst *BST[K, T]) Get(key K) (T, error) {
	if bst.root == nil {
		var res T
		return res, errors.New("the tree is empty")
	}

	isCorrect, node := walk[K, T](key, bst.root)

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

func (bst *BST[K, T]) Delete(key K) (T, error) {
	if bst.root == nil {
		var res T
		return res, errors.New("the tree is empty")
	}

	isCorrect, node := walk[K, T](key, bst.root)

	if !isCorrect {
		var res T
		return res, errors.New("the key could not be found")
	}
	ret := node.value
	bst.root = bst._delete(key, bst.root)
	return ret, nil
}

func (bst *BST[K, T]) _delete(key K, node *node[K, T]) *node[K, T] {
	if node == nil {
		return nil
	}

	var cmp = node.compare(key)
	if cmp < 0 {
		node.left = bst._delete(key, node.left)
	} else if cmp > 0 {
		node.right = bst._delete(key, node.right)
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

func (node *node[K, T]) compare(key K) int8 {
	if node.key < key {
		return 1
	} else if node.key > key {
		return -1
	}
	return 0
}

func (bst *BST[K, T]) Min() (K, error) {
	if bst.root == nil {
		var ret K
		return ret, errors.New("the tree is empty")
	}
	ret := bst.root.Min()
	return ret.key, nil
}

func (node *node[K, T]) Min() *node[K, T] {
	curr := node
	next := curr

	for next != nil {
		curr = next
		next = next.left
	}
	return curr
}

func (bst *BST[K, T]) Max() (K, error) {
	if bst.root == nil {
		var ret K
		return ret, errors.New("the tree is empty")
	}
	ret := bst.root.Max()
	return ret.key, nil
}

func (node *node[K, T]) Max() *node[K, T] {
	curr := node
	next := curr

	for next != nil {
		curr = next
		next = next.right
	}
	return curr
}

func (bst *BST[K, T]) DeleteMax() {
	if bst.root == nil {
		panic("the tree is empty")
	}
	bst.root = deleteMax(bst.root)
}

func deleteMax[K constraints.Ordered, T any](node *node[K, T]) *node[K, T] {
	if node.right == nil {
		return node.left
	}
	node.right = deleteMax(node.right)
	return node
}

func (bst *BST[K, T]) DeleteMin() {
	if bst.root == nil {
		panic("the tree is empty")
	}
	bst.root = deleteMin(bst.root)
}

func deleteMin[K constraints.Ordered, T any](node *node[K, T]) *node[K, T] {
	if node.left == nil {
		return node.right
	}
	node.left = deleteMin(node.left)
	return node
}

func (bst *BST[K, T]) Floor(key K) (K, error) {
	if bst.root == nil {
		var ret K
		return ret, errors.New("the tree is empty")
	}

	ret := floor[K, T](key, bst.root)
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
	cmp := node.compare(key) * -1

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

func (bst *BST[K, T]) Ceiling(key K) (K, error) {
	if bst.root == nil {
		var ret K
		return ret, errors.New("the tree is empty")
	}
	ret := ceiling[K, T](key, bst.root)
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
	cmp := node.compare(key) * -1

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

func (bst *BST[K, T]) GetAllKeys() ([]K, error) {
	if bst.root == nil {
		return nil, errors.New("the tree is empty")
	}
	minVal, err := bst.Min()
	maxVal, err2 := bst.Max()
	if err != nil {
		return nil, err
	}
	if err2 != nil {
		return nil, err2
	}
	ret, _ := bst.GetKeysInRange(minVal, maxVal)
	return ret, nil
}

func (bst *BST[K, T]) GetKeysInRange(lo K, hi K) ([]K, error) {
	if bst.root == nil {
		return nil, errors.New("the tree is empty")
	}
	if lo > hi {
		return nil, errors.New("the low key is greater than the high key")
	}
	var arr []K
	var tmp []T
	keys(bst.root, &arr, &tmp, lo, hi)
	return arr, nil
}

func keys[K constraints.Ordered, T any](node *node[K, T], listKeys *[]K, listValues *[]T, lo K, hi K) {
	if node == nil {
		return
	}
	if lo < node.key {
		keys(node.left, listKeys, listValues, lo, hi)
	}
	if lo <= node.key && hi >= node.key {
		*listKeys = append(*listKeys, node.key)
		*listValues = append(*listValues, node.value)
	}
	if hi > node.key {
		keys(node.right, listKeys, listValues, lo, hi)
	}
}

func (bst *BST[K, T]) IsEmpty() bool {
	return bst.root == nil
}

func (bst *BST[K, T]) GetAllValues() ([]T, error) {
	if bst.root == nil {
		return nil, errors.New("the tree is empty")
	}
	minVal, err := bst.Min()
	maxVal, err2 := bst.Max()
	if err != nil {
		return nil, err
	}
	if err2 != nil {
		return nil, err2
	}
	var list []T
	var tmp []K
	keys(bst.root, &tmp, &list, minVal, maxVal)
	return list, nil
}

func (bst *BST[K, T]) Mapify() (map[K]T, error) {
	if bst.root == nil {
		return nil, errors.New("the tree is empty")
	}

	minVal, err := bst.Min()
	maxVal, err2 := bst.Max()
	if err != nil {
		return nil, err
	}
	if err2 != nil {
		return nil, err2
	}

	var allKeys []K
	var allValues []T
	keys(bst.root, &allKeys, &allValues, minVal, maxVal)
	ret := make(map[K]T)

	for i := 0; i < len(allKeys); i++ {
		ret[allKeys[i]] = allValues[i]
	}

	return ret, nil
}
