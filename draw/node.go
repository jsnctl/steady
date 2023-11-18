package draw

type Node struct {
	Name     string
	Root     *Node
	Children []*Node
}

func (n *Node) AddChild(child *Node) {
	if n.Root == nil { // case where node is root
		n.Root = n
		child.Root = n
	} else {
		child.Root = n.Root
	}

	if !IsNodeInTree(n.Root, child) {
		n.Children = append(n.Children, child)
	}
}

func IsNodeInTree(root *Node, node *Node) bool {
	if root.Name == node.Name {
		return true
	}

	for _, child := range root.Children {
		if IsNodeInTree(child, node) {
			return true
		}
	}

	return false
}
