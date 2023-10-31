package draw

import "fmt"

type Tree struct {
	Nodes []*Node
}

func (t *Tree) AddNode(node *Node) {
	t.Nodes = append(t.Nodes, node)
}

func (t *Tree) Draw() {
	for _, n := range t.Nodes {
		fmt.Print(n.Name + "->(")

		for _, c := range n.Children {
			fmt.Print(c.Name + ",")
		}
		fmt.Print(")\n")
	}
}
