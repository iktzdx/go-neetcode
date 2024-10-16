package serializeanddeserializebinarytree_test

import (
	"reflect"
	"testing"

	"github.com/iktzdx/go-neetcode/trees"
	solution "github.com/iktzdx/go-neetcode/trees/serialize-and-deserialize-binary-tree"
)

func Test_Codec(t *testing.T) {
	tests := map[string]struct {
		root []int
	}{
		"case #1": {
			root: []int{},
		},
		"case #2": {
			root: []int{1, 2, 3, trees.NULL, trees.NULL, 4, 5},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			root := trees.BSTFromArrayInts(test.root)

			set := solution.Constructor()
			serialized := set.Serialize(root)

			deser := solution.Constructor()
			deserialized := deser.Deserialize(serialized)

			got := trees.BSTToArrayInts(deserialized)

			if !reflect.DeepEqual(test.root, got) {
				t.Fatalf("(De/S)erialize(%v): expected = %v, got = %v", test.root, test.root, got)
			}
		})
	}
}
