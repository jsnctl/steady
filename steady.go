package main

import (
	"flag"
	"fmt"
	"github.com/jsnctl/steady/draw"
	"helm.sh/helm/v3/pkg/chart/loader"
	"strings"
)

const DefaultChart = "charts/example"

func main() {
	valuesFile := flag.String("chart", DefaultChart, "Select chart directory")
	flag.Parse()

	chart, err := loader.Load(*valuesFile)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	root := draw.Node{Name: "root"}
	for _, f := range chart.Raw {
		split := strings.Split(f.Name, "/")
		node := draw.Node{Name: split[0]}
		root.AddChild(&node)
		if len(split) > 0 {
			for _, s := range split[1:] {
				child := draw.Node{Name: s}
				node.AddChild(&child)
			}
		}
	}
	fmt.Println("!!!")
}
