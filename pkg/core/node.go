package core

type Node struct {
	Name          string
	FormattedName string
	IsRoot        bool
	IsYaml        bool
	IsValuesFile  bool
	Data          []uint8
	Children      []*Node
}

type ChartFileMeta struct {
	Path []string
	Name string
	Data []uint8
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
