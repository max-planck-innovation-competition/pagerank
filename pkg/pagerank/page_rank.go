package pagerank

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// PageRank is a struct that contains the page rank implementation
type PageRank struct {
	Alpha     float64 // Alpha is the damping parameter for PageRank, default=0.85.
	MaxIter   uint    // MaxIter is the max amount of iterations
	Tolerance float64 // Tolerance
	N         int     // N is the amount of nodes in the graph
	Iteration uint    // Iteration is the counter of iterations of the implementation
	sortIndex []NodeID
	*Graph
}

// NewPageRank creates a new PageRank implementation.
func NewPageRank(g *Graph) *PageRank {
	pr := PageRank{
		Alpha:     0.85,
		MaxIter:   1000,
		Tolerance: 1e-12,
	}
	pr.Graph = g
	return &pr
}

// CalcPageRank calculates the PageRank of the graph
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
		fmt.Println("--------------------------------------------")
		fmt.Println("\t Iteration: ", pr.Iteration, "/", pr.MaxIter)
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
				// fmt.Println(nodeKey, edge.From.Id, "\t\t", cursorNode.Rank, "/", float64(cursorNode.OutDegree()), "=", currentRankDividedByOutDegree)
				// add the result to the sum
				rankSumOfConnectedNode += currentRankDividedByOutDegree
				// fmt.Println(nodeKey, edge.From.Id, rankSumOfConnectedNode)
			}
			newRank := 1/float64(pr.N) + pr.Alpha*rankSumOfConnectedNode
			//fmt.Println(nodeKey, "new rank", newRank)
			// compute the L1 norm
			absErr := math.Abs(newRank - pr.Nodes[nodeKey].Rank)
			// fmt.Println(nodeKey, "absErr", newRank, "-", pr.Nodes[nodeKey].Rank, "=", absErr)
			l1Err += absErr
			// set the new rank value to the current node
			pr.Nodes[nodeKey].Rank = newRank
		}
		fmt.Println("\t new err", l1Err)
		// check if the L1 error is smaller than the initial tolerance
		// fmt.Println(l1Err, pr.Tolerance, l1Err < pr.Tolerance)
		if l1Err < pr.Tolerance {
			// fmt.Println("terminate")
			// norm the results by dividing by the sum
			rankSum := pr.SumTotalNodeRank()
			// fmt.Println("terminate")
			for _, node := range pr.Nodes {
				node.Rank = node.Rank / rankSum
			}
			return
		}
	}
	panic("PageRank did not converge in " + strconv.Itoa(int(pr.MaxIter)) + " iterations.")
}

// RemoveSelfLoops removes self loops from the graph.
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
	res := "Page Rank \n"
	res += "Nodes: " + strconv.Itoa(len(pr.Nodes)) + "\n"
	res += "Edges: " + strconv.Itoa(len(pr.Edges)) + "\n"
	res += "Iterations: " + strconv.Itoa(int(pr.Iteration)) + "\n"
	res += "----------------------------------\n"
	for _, n := range pr.Nodes {
		rank := fmt.Sprintf("%f", n.Rank)
		res += string(n.Id) + "\t\t\t" + rank + "\n"
	}
	res += "----------------------------------\n"
	return res
}

// Len returns the length of the sortIndex
func (pr *PageRank) Len() int {
	pr.sortIndex = make([]NodeID, len(pr.Nodes))
	i := 0
	for _, k := range pr.Nodes {
		pr.sortIndex[i] = k.Id
		i++
	}
	return i
}

// Less orders the sortIndex by the rank of the nodes
func (pr *PageRank) Less(i, j int) bool {
	// use the sort map to resolve the index to the node
	firstNodeID := pr.sortIndex[i]
	secondNodeID := pr.sortIndex[j]
	return pr.Nodes[firstNodeID].Rank < pr.Nodes[secondNodeID].Rank
}

// Swap swaps the sortIndex
func (pr *PageRank) Swap(i, j int) {
	pr.sortIndex[i], pr.sortIndex[j] = pr.sortIndex[j], pr.sortIndex[i]
	return
}

// OrderResults orders the results by the rank of the nodes
func (pr *PageRank) OrderResults() {
	fmt.Println("Ordering results")
	sort.Sort(pr)
}

// GetMinToMaxOrder returns the order from min to max
func (pr *PageRank) GetMinToMaxOrder() []NodeID {
	if len(pr.sortIndex) == 0 {
		pr.OrderResults()
	}
	return pr.sortIndex
}

// GetMaxToMinOrder returns the order from max to min
func (pr *PageRank) GetMaxToMinOrder() []NodeID {
	if len(pr.sortIndex) == 0 {
		pr.OrderResults()
	}
	s := make([]NodeID, len(pr.sortIndex))
	copy(s, pr.sortIndex)
	// reverse order
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
