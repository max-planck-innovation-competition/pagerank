package pagerank

import (
	"fmt"
	"testing"
)

var g = NewGraph()

func init() {
	g.AddEdge("0", "1").
		AddEdge("0", "2").
		AddEdge("1", "0").
		AddEdge("1", "2").
		AddEdge("2", "0").
		AddEdge("2", "3").
		AddEdge("3", "0")
}

func TestNewPageRank(t *testing.T) {
	pr := NewPageRank(g)
	pr.CalcPageRank()
	fmt.Println(pr)
}

func TestSortPageRank(t *testing.T) {
	pr := NewPageRank(g)
	pr.CalcPageRank()
	pr.OrderResults()

	fmt.Println("Max to Min")
	for _, k := range pr.GetMaxToMinOrder() {
		fmt.Println("ID:", k, "\t\tRank:", pr.Nodes[k].Rank)
	}

	fmt.Println("Min to Max")
	for _, k := range pr.GetMinToMaxOrder() {
		fmt.Println("ID:", k, "\t\tRank:", pr.Nodes[k].Rank)
	}

}
