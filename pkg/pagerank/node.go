package pagerank

type NodeID string

// OutDegree is the number of neighbors of a node based on outgoing edges
func (n *NodeID) String() string {
	return n.String()
}

type Node struct {
	Id       NodeID
	Rank     float64
	Outgoing map[EdgeID]*Edge
	Incoming map[EdgeID]*Edge
}

func NewNode(id string) *Node {
	node := Node{
		Id:       NodeID(id),
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
