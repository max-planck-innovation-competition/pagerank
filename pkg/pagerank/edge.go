package pagerank

// EdgeID is a unique identifier for an edge
type EdgeID string

// Edge represents an edge in the graph
type Edge struct {
	Id     EdgeID
	From   *Node
	To     *Node
	Weight float64
}

// GenerateEdgeID generates a new edge ID
func GenerateEdgeID(from, to *Node) EdgeID {
	return generateEdgeIDFromNodeIDs(from.Id, to.Id)
}

// generateEdgeIDFromNodeIDs generates a new edge ID
func generateEdgeIDFromNodeIDs(from, to NodeID) EdgeID {
	return EdgeID(from + ":" + to)
}

// NewEdge creates a new edge
func NewEdge(from, to *Node, weight float64) *Edge {
	edgeID := GenerateEdgeID(from, to)
	edge := Edge{
		Id:     edgeID,
		From:   from,
		To:     to,
		Weight: weight,
	}
	// create links
	from.Outgoing[edgeID] = &edge
	to.Incoming[edgeID] = &edge
	return &edge
}
