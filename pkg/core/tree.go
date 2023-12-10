package core

import (
	"fmt"
	"strings"
)

func NewTree() *Node {
	return &Node{Name: RootName, IsRoot: true}
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

type ValuesData struct {
	Image struct {
		Repository string `yaml:"repository"`
	} `yaml:"image"`
}

func GetRootValuesNode(tree *Node) *Node {
	for _, child := range tree.Children {
		if child.Name == "values.yaml" {
			return child
		}
	}
	return nil
}
