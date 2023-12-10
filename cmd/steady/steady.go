package main

import (
	"fmt"
	"github.com/jsnctl/steady/pkg/core"
	"gopkg.in/yaml.v2"
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

	var flatTree []core.ChartFileMeta
	for _, f := range chart.Raw {
		split := strings.Split(f.Name, "/")
		chartFileMeta := core.ChartFileMeta{
			Path: split,
			Name: f.Name,
			Data: f.Data,
		}
		flatTree = append(flatTree, chartFileMeta)
	}

	fmt.Println("Steady v" + Version)
	fmt.Println("")

	tree := core.FlatTreeToTree(flatTree)
	core.Draw(tree, 0)
	fmt.Println("")

	rootValues := core.GetRootValuesNode(tree)
	var valuesData core.ValuesData
	err = yaml.Unmarshal(rootValues.Data, &valuesData)
	if err != nil {
		panic(err)
	}

	fmt.Println("Images (from ./values.yaml):")
	formattedImage := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 95, valuesData.Image.Repository)
	fmt.Println("- " + formattedImage)
}
