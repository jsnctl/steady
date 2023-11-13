package draw

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode(t *testing.T) {
	A := Node{Name: "A"}
	B := Node{Name: "B"}
	C := Node{Name: "C"}
	D := Node{Name: "D"}
	E := Node{Name: "E"}

	A.AddChild(&B)
	A.AddChild(&C)
	C.AddChild(&D)

	assert.Equal(t, 2, len(A.Children))
	assert.Equal(t, 1, len(C.Children))

	assert.True(t, IsNodeInTree(&A, &D))
	assert.True(t, IsNodeInTree(&A, &C))
	assert.False(t, IsNodeInTree(&A, &E))
	assert.False(t, IsNodeInTree(&C, &A))
}
