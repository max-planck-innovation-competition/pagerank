package pagerank

import (
	"fmt"
	"testing"
)

func TestNewPageRank(t *testing.T) {
	g := NewGraph()
	g.AddEdge("0", "1").
		AddEdge("0", "2").
		AddEdge("1", "0").
		AddEdge("1", "2").
		AddEdge("2", "0").
		AddEdge("2", "3").
		AddEdge("3", "0")
	pr := NewPageRank(g)
	pr.CalcPageRank()
	fmt.Println(pr)
}
