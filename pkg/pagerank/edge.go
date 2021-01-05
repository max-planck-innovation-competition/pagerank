package pagerank

type EdgeID string

type Edge struct {
	Id     EdgeID
	From   *Node
	To     *Node
	Weight float64
}

func GenerateEdgeID(from, to *Node) EdgeID {
	return EdgeID(from.Id + ":" + to.Id)
}

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
