package main

import (
	"flag"
	"fmt"
)

const DefaultChart = "charts/example"

func main() {
	cmd := flag.String("chart", DefaultChart, "Select chart directory")
	flag.Parse()
	fmt.Printf(*cmd)
}
