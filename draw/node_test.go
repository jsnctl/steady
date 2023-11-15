package draw

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Tree struct {
	nodes map[string]*Node
}

/*
		 A
		/ \
	   B   C
			\
			 D
*/
func testTree() *Tree {
	nodeMap := make(map[string]*Node)
	nodes := []string{"A", "B", "C", "D", "E"}
	for _, node := range nodes {
		nodeMap[node] = &Node{Name: node}
	}

	tree := &Tree{nodes: nodeMap}

	tree.nodes["A"].AddChild(tree.nodes["B"])
	tree.nodes["A"].AddChild(tree.nodes["C"])
	tree.nodes["C"].AddChild(tree.nodes["D"])

	return tree
}
func TestNode(t *testing.T) {
	tree := testTree()

	assert.Equal(t, 2, len(tree.nodes["A"].Children))
	assert.Equal(t, 1, len(tree.nodes["C"].Children))

	assert.True(t, IsNodeInTree(tree.nodes["A"], tree.nodes["D"]))
	assert.True(t, IsNodeInTree(tree.nodes["A"], tree.nodes["C"]))
	assert.False(t, IsNodeInTree(tree.nodes["A"], tree.nodes["E"]))
	assert.False(t, IsNodeInTree(tree.nodes["C"], tree.nodes["A"]))
}
