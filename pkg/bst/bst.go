package bst

// btree.go
type CustomInt int

func (i CustomInt) Equal(a Comparable) bool {
	return i == a.(CustomInt)
}

func (i CustomInt) Less(b Comparable) bool {
	return i < b.(CustomInt)
}

type CustomString string

func (i CustomString) Equal(a Comparable) bool {
	return i == a.(CustomString)
}

func (i CustomString) Less(b Comparable) bool {
	return i < b.(CustomString)
}

type CustomFloat float64

func (i CustomFloat) Equal(a Comparable) bool {
	return i == a.(CustomFloat)
}

func (i CustomFloat) Less(b Comparable) bool {
	return i < b.(CustomFloat)
}

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

type BSTStringString struct {
	Root *Node[CustomString, CustomString]
}

// Contains implements SymbolTable.
func (b *BST[K, V]) Contains(key K) bool {
	bstKeys := b.Keys()
	for i := range bstKeys {
		if bstKeys[i].Equal(key) {
			return true
		}
	}
	return false
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
	if b.IsEmpty() {
		b.Root = &Node[K, V]{Key: key, Value: value}
		return
	}

	currentNode := b.Root
	for {
		// Key exists, update the value
		if currentNode.Key.Equal(key) {
			currentNode.Value = value
			return
		}

		if !currentNode.Key.Less(key) {
			// if node left is null, create a new node and return
			if currentNode.Left == nil {
				currentNode.Left = &Node[K, V]{Key: key, Value: value}
				return
			}

			// visit left child
			currentNode = currentNode.Left
			continue
		}

		if currentNode.Key.Less(key) {
			// if node right is null, create a new node and return
			if currentNode.Right == nil {
				currentNode.Right = &Node[K, V]{Key: key, Value: value}
				return
			}

			// visit right child
			currentNode = currentNode.Right
			continue
		}
	}
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

// PreorderIterKeys will traverse the keys in the tree by first visiting the root node, then the left node and then the right node
// to do this iteratively we will use a stack. we will add items on top of the stack and remove them from the top of the stack
// when the stack is empty, then we have finished.
// on every iteration we will
// 1. take the item from the top of the stack. we will add the key to the response array
// 2. add the left node to the stack, add the right item to the stack
func (b *BSTStringString) preOrderIterKeys() []CustomString {
	if b.Root == nil {
		return nil
	}

	res := []CustomString{}

	type LinkedListNode struct {
		Node *Node[CustomString, CustomString]
		Next *LinkedListNode
	}

	stack := &LinkedListNode{
		Node: b.Root,
	}

	for stack != nil {
		currentNode := stack.Node
		// move head to pop from stack
		stack = stack.Next

		// add root to response
		res = append(res, currentNode.Key)

		// add right node
		if currentNode.Right != nil {
			newhead := &LinkedListNode{
				Node: currentNode.Right,
				Next: stack,
			}
			stack = newhead
		}

		// add left node
		if currentNode.Left != nil {
			newhead := &LinkedListNode{
				Node: currentNode.Left,
				Next: stack,
			}
			stack = newhead
		}
	}

	return res
}

func (b *BSTStringString) inorderKeysIter() []CustomString {
	if b.Root == nil {
		return nil
	}

	type LinkedListNode struct {
		Node *Node[CustomString, CustomString]
		Next *LinkedListNode
	}

	var stack *LinkedListNode
	current := b.Root

	res := []CustomString{}
	for current != nil || stack != nil {
		// Go as far left as possible, pushing nodes
		for current != nil {
			stack = &LinkedListNode{Node: current, Next: stack}
			current = current.Left
		}

		// Pop
		poppedNode := stack.Node
		stack = stack.Next

		// Add to response
		res = append(res, poppedNode.Key)

		// Explore right subtree next
		current = poppedNode.Right
	}

	return res
}
