package bst

import (
	"reflect"
	"testing"
)

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

func TestCreateBSTTree(t *testing.T) {
	// Create a new BST
	bst := &BST[CustomInt, string]{
		Root: &Node[CustomInt, string]{
			Key:   10,
			Value: "root",
		},
	}
	// Check if the BST is empty
	if bst.IsEmpty() {
		t.Errorf("Expected BST to be non-empty")
	}
}

func TestGetBST(t *testing.T) {
	// Create a new BST
	bst := &BST[CustomInt, string]{
		Root: &Node[CustomInt, string]{
			Key:   10,
			Value: "root",
			Left: &Node[CustomInt, string]{
				Key:   5,
				Value: "left",
				Left: &Node[CustomInt, string]{
					Key:   3,
					Value: "left.left",
				},
			},
			Right: &Node[CustomInt, string]{
				Key:   15,
				Value: "right",
				Left: &Node[CustomInt, string]{
					Key:   12,
					Value: "right.left",
				},
				Right: &Node[CustomInt, string]{
					Key:   18,
					Value: "right.right",
				},
			},
		},
	}
	// Get the value for an existing key
	value, ok := bst.Get(10)
	if !ok || value != "root" {
		t.Errorf("Expected to get 'root' for key 10")
	}
	// Get the value for a non-existing key
	value, ok = bst.Get(20)
	if ok {
		t.Errorf("Expected to not find a value for key 20")
	}

	value, ok = bst.Get(10)
	if !ok {
		t.Errorf("Expected to get 'root' for key 10")
	}
	if value != "root" {
		t.Errorf("Expected to get 'root' for key 10")
	}

	value, ok = bst.Get(18)
	if !ok {
		t.Errorf("Expected to get 'right.right' for key 18")
	}
	if value != "right.right" {
		t.Errorf("Expected to get 'right.right' for key 18")
	}
}

func TestKeys(t *testing.T) {
	bst := &BST[CustomInt, string]{
		Root: &Node[CustomInt, string]{
			Key:   10,
			Value: "root",
			Left: &Node[CustomInt, string]{
				Key:   5,
				Value: "left",
				Left: &Node[CustomInt, string]{
					Key:   3,
					Value: "left.left",
				},
			},
			Right: &Node[CustomInt, string]{
				Key:   15,
				Value: "right",
				Left: &Node[CustomInt, string]{
					Key:   12,
					Value: "right.left",
				},
				Right: &Node[CustomInt, string]{
					Key:   18,
					Value: "right.right",
				},
			},
		},
	}

	keys := bst.Keys()
	expectedKeys := []CustomInt{10, 5, 15, 3, 12, 18}

	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Errorf("Expected keys %v, but got %v", expectedKeys, keys)
	}

}
