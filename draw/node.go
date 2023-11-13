package draw

type Node struct {
	Name     string
	Root     *Node
	Children []*Node
}

func (n *Node) AddChild(child *Node) {
	if n.Root == nil {
		child.Root = n
		n.Children = append(n.Children, child)
		return
	}

	if !IsNodeInTree(n.Root, child) {
		child.Root = n.Root
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
