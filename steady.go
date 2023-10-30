package main

import (
	"flag"
	"fmt"
	"helm.sh/helm/v3/pkg/chart/loader"
)

const DefaultChart = "charts/example"

func main() {
	valuesFile := flag.String("chart", DefaultChart, "Select chart directory")
	flag.Parse()

	chart, err := loader.Load(*valuesFile)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	for _, f := range chart.Raw {
		fmt.Printf("    |    \n")
		fmt.Printf("|" + f.Name + "|\n")
		fmt.Printf("    |    \n")
	}
}
