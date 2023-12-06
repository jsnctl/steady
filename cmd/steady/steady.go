package main

import (
	"fmt"
	"github.com/jsnctl/steady/pkg/core"
	"helm.sh/helm/v3/pkg/chart/loader"
	"strings"
)

const DefaultChart = "charts/example"
const Version = "0.0.1"

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

	fmt.Println("Steady v" + Version)
	fmt.Println("")

	tree := core.FlatTreeToTree(flatTree)
	core.Draw(tree, 0)
	fmt.Println("")
}
