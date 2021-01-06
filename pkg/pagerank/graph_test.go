package pagerank

import (
	"fmt"
	"testing"
)

func TestNewGraph(t *testing.T) {

	g := NewGraph()
	g.AddEdge("home", "imprint").
		AddEdge("home", "agb").
		AddEdge("imprint", "agb").
		AddEdge("agb", "privacy").
		AddEdge("privacy", "home")

	fmt.Println(g)
}
