package pagerank

import (
	"fmt"
	"testing"
)

func TestGenerateGraphFromCSV(t *testing.T) {
	g, err := GenerateGraphFromCSV("./test-data/test.csv", true, 0, 1)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("edges", g.GetAmountOfEdges())
	fmt.Println("nodes", g.GetAmountOfNodes())
	pr := NewPageRank(g)
	pr.CalcPageRank()
	fmt.Println("PageRank", pr.GetMaxToMinOrder())
}
