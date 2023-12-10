package core

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

}

func TestAddChild_IsIdempotent(t *testing.T) {
	tree := testTree()

	for i := 0; i < 10; i++ {
		tree.nodes["A"].AddChild(tree.nodes["B"])
	}

	assert.Equal(t, 2, len(tree.nodes["A"].Children))
}

func TestAddChild_MultipleChildrenOfNonRoot(t *testing.T) {
	tree := testTree()
	assert.Equal(t, 1, len(tree.nodes["C"].Children))

	/*
			 A
			/ \
		   B   C
			  /	\
			 D	 E
	*/
	tree.nodes["C"].AddChild(tree.nodes["E"])
	assert.Equal(t, 2, len(tree.nodes["C"].Children))
}

func TestFlatTreeToTree(t *testing.T) {
	flatTree := []ChartFileMeta{
		{Path: []string{"A", "B"}},
		{Path: []string{"A", "C", "D"}},
		{Path: []string{"A", "C", "E"}},
	}
	tree := FlatTreeToTree(flatTree)
	assert.Equal(t, 2, len(tree.Children[0].Children))             // B, C
	assert.Equal(t, 2, len(tree.Children[0].Children[1].Children)) // D, E

	/*
				 A				K
				/|\			   / \
			   B C D		  L	  M
			  /	  /	\
			 E	 F	 G
					/|\
		           H I J
	*/
	deeperFlatTree := []ChartFileMeta{
		{Path: []string{"A", "B", "E"}},
		{Path: []string{"A", "C"}},
		{Path: []string{"A", "D", "F"}},
		{Path: []string{"A", "D", "G", "H"}},
		{Path: []string{"A", "D", "G", "I"}},
		{Path: []string{"A", "D", "G", "J"}},
		{Path: []string{"K", "L"}},
		{Path: []string{"K", "M"}},
	}
	tree = FlatTreeToTree(deeperFlatTree)
	assert.Equal(t, 2, len(tree.Children))                                     // A, K
	assert.Equal(t, 3, len(tree.Children[0].Children))                         // B, C, D
	assert.Equal(t, 3, len(tree.Children[0].Children[2].Children[1].Children)) // H, I, J

}
