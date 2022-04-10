package pagerank

import (
	"fmt"
	"testing"
)

func TestGenerateGraphFromCSV(t *testing.T) {
	g := NewGraph()
	err := g.AddNodesFromCSV("./test-data/test.csv", true, 0, 1)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("edges", g.GetAmountOfEdges())
	fmt.Println("nodes", g.GetAmountOfNodes())
	pr := NewPageRank(g)
	pr.CalcPageRank()
	pr.OrderResults()
	fmt.Println("Max to Min")
	for i, k := range pr.GetMaxToMinOrder() {
		fmt.Println("ID:", k, "\t\tRank:", pr.Nodes[k].Rank)
		if i > 1000 {
			break
		}
	}

}
