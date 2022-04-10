package pagerank

import (
	"fmt"
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph().
		AddEdge("home", "imprint").
		AddEdge("home", "agb").
		AddEdge("imprint", "agb").
		AddEdge("agb", "privacy").
		AddEdge("privacy", "home")
	fmt.Println(g)
}

func TestGraph_RemoveEdge(t *testing.T) {
	g := NewGraph().
		AddEdge("0", "1").
		AddEdge("0", "2").
		AddEdge("1", "0").
		AddEdge("1", "2").
		AddEdge("2", "0").
		AddEdge("2", "3").
		AddEdge("3", "0")

	if g.GetAmountOfEdges() != 7 {
		t.Error("wrong amount of edges:", g.GetAmountOfEdges())
	}

	g.AddEdge("4", "5")

	if g.GetAmountOfEdges() != 8 {
		t.Error("wrong amount of edges:", g.GetAmountOfEdges())
	}

	if !g.CheckIfNodeExists("5") {
		t.Error("Node 5 should exist")
	}
	// remove edge
	g.RemoveEdge("4", "5")

	// check if node 5 is still there
	if !g.CheckIfNodeExists("5") {
		t.Error("Node 5 should exist")
	}
	// check if node 4 is still there
	if !g.CheckIfNodeExists("4") {
		t.Error("Node 4 should exist")
	}
	// check edges
	if g.CheckIfEdgeExistFromNodes("4", "5") {
		t.Error("edge from 4 to 5 should not exist")
	}
	// check edges
	if g.CheckIfEdgeExistFromNodes("5", "4") {
		t.Error("edge from 5 to 4 should not exist")
	}
}

func TestGraph_RemoveNode(t *testing.T) {
	g := NewGraph().
		AddEdge("0", "1").
		AddEdge("0", "2").
		AddEdge("1", "0").
		AddEdge("1", "2").
		AddEdge("2", "0").
		AddEdge("2", "3").
		AddEdge("3", "0")

	if g.GetAmountOfEdges() != 7 {
		t.Error("wrong amount of edges:", g.GetAmountOfEdges())
	}

	if g.GetAmountOfNodes() != 4 {
		t.Error("wrong amount of nodes:", g.GetAmountOfNodes())
	}

	g.RemoveNode("0")

	if g.GetAmountOfNodes() != 3 {
		t.Error("wrong amount of nodes:", g.GetAmountOfNodes())
	}

	// check nodes
	if g.CheckIfNodeExists("0") {
		t.Error("Node 0 should not exist")
	}
	// check edges
	if g.CheckIfEdgeExistFromNodes("0", "1") {
		t.Error("edge from 0 to 1 should not exist")
	}
	if g.CheckIfEdgeExistFromNodes("0", "2") {
		t.Error("edge from 0 to 2 should not exist")
	}
	if g.CheckIfEdgeExistFromNodes("1", "0") {
		t.Error("edge from 1 to 0 should not exist")
	}
	if !g.CheckIfEdgeExistFromNodes("1", "2") {
		t.Error("edge from 1 to 2 should exist")
	}

}
