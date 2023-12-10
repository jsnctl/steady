package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testNodes map[string]*Node

/*
		 A
		/ \
	   B   C
			\
			 D
*/
func initTestNodes() {
	nodeMap := make(map[string]*Node)
	nodes := []string{"A", "B", "C", "D", "E"}
	for _, node := range nodes {
		nodeMap[node] = &Node{Name: node}
	}
	testNodes = nodeMap

	testNodes["A"].AddChild(testNodes["B"])
	testNodes["A"].AddChild(testNodes["C"])
	testNodes["C"].AddChild(testNodes["D"])
}

func TestNode(t *testing.T) {
	initTestNodes()

	assert.Equal(t, 2, len(testNodes["A"].Children))
	assert.Equal(t, 1, len(testNodes["C"].Children))

}

func TestAddChild_IsIdempotent(t *testing.T) {
	initTestNodes()

	for i := 0; i < 10; i++ {
		testNodes["A"].AddChild(testNodes["B"])
	}

	assert.Equal(t, 2, len(testNodes["A"].Children))
}

func TestAddChild_MultipleChildrenOfNonRoot(t *testing.T) {
	initTestNodes()
	assert.Equal(t, 1, len(testNodes["C"].Children))

	/*
			 A
			/ \
		   B   C
			  /	\
			 D	 E
	*/
	testNodes["C"].AddChild(testNodes["E"])
	assert.Equal(t, 2, len(testNodes["C"].Children))
}
