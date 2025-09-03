package bst

import (
	"reflect"
	"testing"
)

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

func TestContains(t *testing.T) {
	b := &BST[CustomInt, string]{
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

	if !b.Contains(10) {
		t.Errorf("Expected to find key 10")
	}
	if b.Contains(20) {
		t.Errorf("Expected to not find key 20")
	}

}

func TestPut(t *testing.T) {
	b := &BST[CustomInt, string]{
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

	b.Put(20, "right.right.right")
	if !b.Contains(20) {
		t.Errorf("Expected to find key 20")
	}

	// Check if the new key is present
	if b.Contains(21) {
		t.Errorf("Expected to not find key 21")
	}
}

func TestBSTStringString_preOrderIterKeys(t *testing.T) {
	type fields struct {
		Root *Node[CustomString, CustomString]
	}
	tests := []struct {
		name   string
		fields fields
		want   []CustomString
	}{
		{
			name: "nil root",
			want: []CustomString{
				CustomString("10"),
				CustomString("5"),
				CustomString("3"),
				CustomString("7"),
				CustomString("15"),
			},
			fields: fields{
				Root: &Node[CustomString, CustomString]{
					Key:   "10",
					Value: "10",
					Left: &Node[CustomString, CustomString]{
						Key:   "5",
						Value: "5",
						Left: &Node[CustomString, CustomString]{
							Key:   "3",
							Value: "3",
						},
						Right: &Node[CustomString, CustomString]{
							Key:   "7",
							Value: "7",
						},
					},
					Right: &Node[CustomString, CustomString]{
						Key:   "15",
						Value: "15",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BSTStringString{
				Root: tt.fields.Root,
			}
			if got := b.preOrderIterKeys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BSTStringString.preOrderIterKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTStringString_inorderKeysIter(t *testing.T) {
	type fields struct {
		Root *Node[CustomString, CustomString]
	}
	tests := []struct {
		name   string
		fields fields
		want   []CustomString
	}{
		{
			name: "nil root",
			want: []CustomString{
				CustomString("3"),
				CustomString("5"),
				CustomString("7"),
				CustomString("10"),
				CustomString("15"),
			},
			fields: fields{
				Root: &Node[CustomString, CustomString]{
					Key:   "10",
					Value: "10",
					Left: &Node[CustomString, CustomString]{
						Key:   "5",
						Value: "5",
						Left: &Node[CustomString, CustomString]{
							Key:   "3",
							Value: "3",
						},
						Right: &Node[CustomString, CustomString]{
							Key:   "7",
							Value: "7",
						},
					},
					Right: &Node[CustomString, CustomString]{
						Key:   "15",
						Value: "15",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BSTStringString{
				Root: tt.fields.Root,
			}
			if got := b.inorderKeysIter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BSTStringString.inorderKeysIter() = %v, want %v", got, tt.want)
			}
		})
	}
}
