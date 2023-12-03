package draw

import (
	"fmt"
	"strings"
)

type Node struct {
	Name     string
	Children []*Node
}

const RootName = "ROOT"

func (n *Node) AddChild(prospective *Node) {
	for _, child := range n.Children {
		if child.Name == prospective.Name {
			child.Children = append(child.Children, prospective.Children...)
			return
		}
	}
	n.Children = append(n.Children, prospective)
}

func FlatTreeToTree(flatTree [][]string) *Node {
	var flatNodes [][]*Node
	for branchIndex, branch := range flatTree {
		flatNodes = append(flatNodes, make([]*Node, len(branch)))
		for nodeIndex, nodeName := range branch {
			flatNodes[branchIndex][nodeIndex] = &Node{
				Name: nodeName,
			}
		}
	}
	tree := Node{Name: RootName}
	pointer := &tree
	for branchIndex, branch := range flatTree {
		pointer = &tree
		for nodeIndex := range branch {
			pointer.AddChild(flatNodes[branchIndex][nodeIndex])
			newPointer := pointer.Children[len(pointer.Children)-1]
			pointer = newPointer
		}
	}

	return &tree
}

func Draw(node *Node, shift int) {
	if node.Name != RootName {
		if shift <= 1 {
			fmt.Println(strings.Repeat(" ", shift*5), node.Name)
		} else {
			fmt.Println(strings.Repeat(" ", shift*5), "\\ "+node.Name)
		}
	}
	for _, node := range node.Children {
		Draw(node, shift+1)
	}
}
