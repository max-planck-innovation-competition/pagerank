package pagerank

import (
	"fmt"
	"math"
	"strconv"
)

type PageRank struct {
	Alpha     float64 // Alpha is the damping parameter for PageRank, default=0.85.
	MaxIter   uint    // MaxIter is the max amount of iterations
	Tolerance float64 // Tolerance
	N         int     // N is the amount of nodes in the graph
	Iteration uint    // Iteration is the counter of iterations of the implementation
	*Graph
}

func NewPageRank(g *Graph) *PageRank {
	pr := PageRank{
		Alpha:     0.85,
		MaxIter:   1000,
		Tolerance: 1e-12,
	}
	pr.Graph = g
	return &pr
}

func (pr *PageRank) CalcPageRank() {
	// check if edges are empty
	if len(pr.Edges) == 0 {
		return
	}
	// remove self loops
	pr.RemoveSelfLoops()
	// check if there are nodes provides
	if len(pr.Nodes) == 0 {
		return
	}
	// init nodes by setting the rank to 1/n
	pr.InitializeNodes()
	// iterate unit convergence
	for pr.Iteration = uint(0); pr.Iteration <= pr.MaxIter; pr.Iteration++ {
		// init the l1Err error
		l1Err := 0.0
		// iterate over nodes in graph
		for nodeKey := range pr.Nodes {
			// calculate the sum of the surrounding nodes of the current node
			rankSumOfConnectedNode := 0.0
			// iterate over the incoming connections
			for _, edge := range pr.Nodes[nodeKey].Incoming {
				cursorNode := edge.From
				// calculate the new rank based on the current rank and the out-degree
				// (out-degree is the number of neighbors of pr node)
				currentRankDividedByOutDegree := cursorNode.Rank / float64(cursorNode.OutDegree())
				// add the result to the sum
				rankSumOfConnectedNode += currentRankDividedByOutDegree
			}
			newRank := 1/float64(pr.N) + pr.Alpha*rankSumOfConnectedNode
			// compute the L1 norm
			absErr := math.Abs(newRank - pr.Nodes[nodeKey].Rank)
			l1Err += absErr
			// set the new rank value to the current node
			pr.Nodes[nodeKey].Rank = newRank
		}
		// check if the L1 error is smaller than the initial tolerance
		if l1Err < pr.Tolerance {
			// norm the results by dividing by the sum
			rankSum := pr.SumTotalNodeRank()
			for _, node := range pr.Nodes {
				node.Rank = node.Rank / rankSum
			}
			return
		}
	}
	panic("PageRank did not converge in " + strconv.Itoa(int(pr.MaxIter)) + " iterations.")
}

func (pr *PageRank) RemoveSelfLoops() {
	// Remove self loops in place
	// no new allocations
	for _, e := range pr.Edges {
		// if the pointers are not the same add the value to the beginning of the array
		if e.From == e.To {
			delete(pr.Edges, e.Id)
		}
	}
}

// InitializeNodes sets the rank of the nodes to 1/n
func (pr *PageRank) InitializeNodes() {
	pr.N = len(pr.Nodes)
	for k := range pr.Nodes {
		pr.Nodes[k].Rank = 1 / float64(pr.N)
	}
}

// SumTotalNodeRank sums up the rank of all nodes
func (pr *PageRank) SumTotalNodeRank() (value float64) {
	for _, n := range pr.Nodes {
		value += n.Rank
	}
	return
}

// To String method
func (pr *PageRank) String() string {
	res := "Pagerank \n"
	res += "Nodes: " + strconv.Itoa(len(pr.Nodes)) + "\n"
	res += "Iterations: " + strconv.Itoa(int(pr.Iteration)) + "\n"
	res += "----------------------------------\n"
	for _, n := range pr.Nodes {
		rank := fmt.Sprintf("%f", n.Rank)
		res += string(n.Id) + "\t\t\t" + rank + "\n"
	}
	res += "----------------------------------\n"
	return res
}
