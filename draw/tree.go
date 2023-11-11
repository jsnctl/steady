package draw

import (
	"fmt"
	"slices"
)

type Tree struct {
	Nodes []*Node
}

func (t *Tree) checkNodeExists(node *Node) bool {
	if slices.Contains(t.Nodes, node) {
		return true
	}
	return false
}

func (t *Tree) AddNode(node *Node) {
	if !t.checkNodeExists(node) {
		t.Nodes = append(t.Nodes, node)
	}
}

func (t *Tree) Draw() {
	for _, n := range t.Nodes {
		fmt.Println(n.Name)
		for _, c := range n.Children {
			fmt.Print("   \\   ")
			fmt.Println("")
			fmt.Print("    " + c.Name)
		}
		fmt.Print("\n")
	}
}
