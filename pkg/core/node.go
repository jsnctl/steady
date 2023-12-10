package core

import (
	"fmt"
	"strings"
)

type Node struct {
	Name          string
	FormattedName string
	IsRoot        bool
	IsYaml        bool
	Data          []uint8
	Children      []*Node
}

type ChartFileMeta struct {
	Path []string
	Name string
	Data []uint8
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

func FlatTreeToTree(flatTree []ChartFileMeta) *Node {
	var flatNodes [][]*Node
	for branchIndex, branch := range flatTree {
		flatNodes = append(flatNodes, make([]*Node, len(branch.Path)))
		for nodeIndex, nodeName := range branch.Path {
			node := &Node{
				Name:          nodeName,
				FormattedName: nodeName, // pre-formatting default
				Data:          branch.Data,
			}
			if strings.HasSuffix(nodeName, ".yaml") {
				node.IsYaml = true
				node.FormattedName = fmt.Sprintf("\x1b[%dm%s\x1b[0m", 92, node.Name)
			}
			flatNodes[branchIndex][nodeIndex] = node
		}
	}
	tree := NewTree()
	pointer := tree
	for branchIndex, branch := range flatTree {
		pointer = tree
		for nodeIndex := range branch.Path {
			pointer.AddChild(flatNodes[branchIndex][nodeIndex])
			newPointer := pointer.Children[len(pointer.Children)-1]
			pointer = newPointer
		}
	}

	return tree
}
