package main

import (
	"fmt"
	"github.com/jsnctl/steady/draw"
	"helm.sh/helm/v3/pkg/chart/loader"
	"strings"
)

const DefaultChart = "charts/example"

func main() {
	chart, err := loader.Load(DefaultChart)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	var flatTree [][]string
	for _, f := range chart.Raw {
		split := strings.Split(f.Name, "/")
		flatTree = append(flatTree, split)
	}

	tree := draw.FlatTreeToTree(flatTree)

	fmt.Println("")
	draw.Draw(tree, 0)
}
