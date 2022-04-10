package pagerank

import (
	"fmt"
	"testing"
)

var prGraph = NewGraph()

func init() {
	prGraph.AddEdge("0", "1").
		AddEdge("0", "2").
		AddEdge("1", "0").
		AddEdge("1", "2").
		AddEdge("2", "0").
		AddEdge("2", "3").
		AddEdge("3", "0")
}

func TestNewPageRank(t *testing.T) {
	pr := NewPageRank(prGraph)
	pr.CalcPageRank()
	fmt.Println(pr)
}

func TestSortPageRank(t *testing.T) {
	pr := NewPageRank(prGraph)
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

func TestNewPageRankAddEdge(t *testing.T) {
	pr := NewPageRank(prGraph)
	// calc initial
	pr.CalcPageRank()
	fmt.Println(pr)
	// add edge and calc again
	prGraph.AddEdge("3", "4")
	pr.CalcPageRank()
	fmt.Println(pr)
	// add more edges
	prGraph.AddEdge("4", "0")
	prGraph.AddEdge("4", "1")
	pr.CalcPageRank()
	fmt.Println(pr)
}

func TestNewPageRankAddEdgeInitial(t *testing.T) {
	prGraph.AddEdge("3", "4")
	prGraph.AddEdge("4", "0")
	prGraph.AddEdge("4", "1")
	pr := NewPageRank(prGraph)
	pr.CalcPageRank()
	fmt.Println(pr)
}
