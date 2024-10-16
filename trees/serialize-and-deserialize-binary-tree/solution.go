package serializeanddeserializebinarytree

import (
	"strconv"
	"strings"

	"github.com/iktzdx/go-neetcode/trees"
)

const (
	sep        = ","
	EmptyValue = "null"
)

type Codec struct {
	data []string
	idx  int
}

func Constructor() *Codec {
	return &Codec{
		data: make([]string, 0),
		idx:  0,
	}
}

type (
	dfsSerializeFunc   func(*trees.TreeNode)
	dfsDeserializeFunc func() *trees.TreeNode
)

// Serializes a tree to a single string.
func (c *Codec) Serialize(root *trees.TreeNode) string {
	var dfs dfsSerializeFunc

	dfs = func(node *trees.TreeNode) {
		if node == nil {
			c.data = append(c.data, EmptyValue)

			return
		}

		c.data = append(c.data, strconv.Itoa(node.Val))

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	return strings.Join(c.data, sep)
}

// Deserializes your encoded data to tree.
func (c *Codec) Deserialize(data string) *trees.TreeNode {
	var dfs dfsDeserializeFunc

	c.data = strings.Split(data, sep)

	dfs = func() *trees.TreeNode {
		if c.data[c.idx] == EmptyValue {
			c.idx++

			return nil
		}

		val, err := strconv.Atoi(c.data[c.idx])
		if err != nil {
			return nil
		}

		node := &trees.TreeNode{Val: val}
		c.idx++

		node.Left = dfs()
		node.Right = dfs()

		return node
	}

	return dfs()
}
