package core

import "strings"

type Node struct {
	Name     string
	IsRoot   bool
	IsYaml   bool
	Children []*Node
}

const RootName = "ROOT"

func NewTree() *Node {
	return &Node{Name: RootName, IsRoot: true}
}

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
			if strings.HasSuffix(nodeName, ".yaml") {
				flatNodes[branchIndex][nodeIndex].IsYaml = true
			}
		}
	}
	tree := NewTree()
	pointer := tree
	for branchIndex, branch := range flatTree {
		pointer = tree
		for nodeIndex := range branch {
			pointer.AddChild(flatNodes[branchIndex][nodeIndex])
			newPointer := pointer.Children[len(pointer.Children)-1]
			pointer = newPointer
		}
	}

	return tree
}
