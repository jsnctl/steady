package core

import (
	"fmt"
	"strings"
)

const TabSize = 4

func Draw(node *Node, depth int) {
	if node.IsRoot {
		// nothing
	} else {
		writeLine(node.FormattedName, depth)
	}
	for _, node := range node.Children {
		Draw(node, depth+1)
	}
}

func writeLine(text string, depth int) {
	indent := ""
	if depth > 1 {
		text = "⌞ " + text
		indent = strings.Repeat(" ", (depth-1)*TabSize)
	}
	fmt.Println(" | " + indent + text)
}
