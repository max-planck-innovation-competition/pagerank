package pagerank

// NodeID is a unique identifier for a node
type NodeID string

// OutDegree is the number of neighbors of a node based on outgoing edges
func (n *NodeID) String() string {
	return n.String()
}

// Node is a node in the graph
type Node struct {
	Id       NodeID
	Rank     float64
	Outgoing map[EdgeID]*Edge
	Incoming map[EdgeID]*Edge
}

func NewNode(id NodeID) *Node {
	node := Node{
		Id:       id,
		Rank:     0,
		Outgoing: map[EdgeID]*Edge{},
		Incoming: map[EdgeID]*Edge{},
	}
	return &node
}

// OutDegree is the number of neighbors of a node based on outgoing edges
func (n *Node) OutDegree() uint {
	return uint(len(n.Outgoing))
}

// InDegree is the number of neighbors of a node based on incoming edges
func (n *Node) InDegree() uint {
	return uint(len(n.Incoming))
}
