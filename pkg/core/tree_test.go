package core

import (
	"github.com/stretchr/testify/assert"
	"helm.sh/helm/v3/pkg/chart/loader"
	"strings"
	"testing"
)

const DefaultChart = "../../charts/example"

func exampleChartToTree() *Node {
	chart, _ := loader.Load(DefaultChart)

	var flatTree []ChartFileMeta
	for _, f := range chart.Raw {
		split := strings.Split(f.Name, "/")
		chartFileMeta := ChartFileMeta{
			Path: split,
			Name: f.Name,
			Data: f.Data,
		}
		flatTree = append(flatTree, chartFileMeta)
	}

	return FlatTreeToTree(flatTree)
}

func TestFlatTreeToTree(t *testing.T) {
	flatTree := []ChartFileMeta{
		{Path: []string{"A", "B"}},
		{Path: []string{"A", "C", "D"}},
		{Path: []string{"A", "C", "E"}},
	}
	tree := FlatTreeToTree(flatTree)
	assert.Equal(t, 2, len(tree.Children[0].Children))             // B, C
	assert.Equal(t, 2, len(tree.Children[0].Children[1].Children)) // D, E

	/*
				 A				K
				/|\			   / \
			   B C D		  L	  M
			  /	  /	\
			 E	 F	 G
					/|\
		           H I J
	*/
	deeperFlatTree := []ChartFileMeta{
		{Path: []string{"A", "B", "E"}},
		{Path: []string{"A", "C"}},
		{Path: []string{"A", "D", "F"}},
		{Path: []string{"A", "D", "G", "H"}},
		{Path: []string{"A", "D", "G", "I"}},
		{Path: []string{"A", "D", "G", "J"}},
		{Path: []string{"K", "L"}},
		{Path: []string{"K", "M"}},
	}
	tree = FlatTreeToTree(deeperFlatTree)
	assert.Equal(t, 2, len(tree.Children))                                     // A, K
	assert.Equal(t, 3, len(tree.Children[0].Children))                         // B, C, D
	assert.Equal(t, 3, len(tree.Children[0].Children[2].Children[1].Children)) // H, I, J

}

func TestGetRootValuesNode(t *testing.T) {
	tree := exampleChartToTree()

	assert.NotNil(t, GetRootValuesNode(tree))
}
