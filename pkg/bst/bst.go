package bst

// btree.go

type Comparable interface {
	Less(comaparedTo Comparable) bool
	Equal(comparedTo Comparable) bool
}

// SymbolTable is a generic associative container interface.
// K must be comparable to be used as a key in maps/trees.
// V can be any type.
type SymbolTable[K Comparable, V any] interface {
	Put(key K, value V)
	Get(key K) (V, bool)
	Delete(key K)
	Contains(key K) bool
	IsEmpty() bool
	Size() int
	Keys() []K
}

var _ SymbolTable[Comparable, any] = (*BST[Comparable, any])(nil)

type BST[K Comparable, V any] struct {
	Root *Node[K, V]
}

// Contains implements SymbolTable.
func (b *BST[K, V]) Contains(key K) bool {
	panic("unimplemented")
}

// Delete implements SymbolTable.
func (b *BST[K, V]) Delete(key K) {
	panic("unimplemented")
}

// Get implements SymbolTable.
func (b *BST[K, V]) Get(key K) (V, bool) {
	if b.Root == nil {
		return *new(V), false
	}

	// if key is equal to root node, return it

	currentNode := b.Root
	for currentNode != nil {
		if currentNode.Key.Equal(key) {
			return currentNode.Value, true
		}
		// if key is smaller than root node, check left
		if key.Less(currentNode.Key) {
			currentNode = currentNode.Left
			continue
		}

		currentNode = currentNode.Right
	}
	return *new(V), false
}

// IsEmpty implements SymbolTable.
func (b *BST[K, V]) IsEmpty() bool {
	if b.Root == nil {
		return true
	}
	return false
}

// Keys implements SymbolTable.
func (b *BST[K, V]) Keys() []K {
	res := []K{}
	tovisit := []*Node[K, V]{b.Root}
	currentIdx := 0
	for currentIdx < len(tovisit) {
		currentNode := tovisit[currentIdx]
		res = append(res, currentNode.Key)

		if currentNode.Left != nil {
			tovisit = append(tovisit, currentNode.Left)
		}
		if currentNode.Right != nil {
			tovisit = append(tovisit, currentNode.Right)
		}
		currentIdx++
	}

	return res
}

// Put implements SymbolTable.
func (b *BST[K, V]) Put(key K, value V) {
	panic("unimplemented")
}

// Size implements SymbolTable.
func (b *BST[K, V]) Size() int {
	return len(b.Keys())
}

type Node[K Comparable, V any] struct {
	Key   K
	Value V
	// Left contains smaller values
	Left *Node[K, V]
	// Right contains larger values
	Right *Node[K, V]
}
