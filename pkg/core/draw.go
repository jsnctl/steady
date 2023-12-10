package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"strings"
)

const TabSize = 4

func Draw(node *Node, depth int) {
	if node.IsRoot {
		// nothing
	} else {
		writeLine(node.FormattedName, depth)
		if node.IsValuesFile {
			var valuesData ValuesData
			err := yaml.Unmarshal(node.Data, &valuesData)
			if err != nil {
				panic(err)
			}
			formattedImage := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 95, valuesData.Image.Repository)
			writeLine(formattedImage, depth+1)
		}
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
