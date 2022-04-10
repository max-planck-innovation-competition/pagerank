package pagerank

// Graph is the data structure of the graph
type Graph struct {
	Nodes map[NodeID]*Node
	Edges map[EdgeID]*Edge
}

// NewGraph initializes a new graph
func NewGraph() (g *Graph) {
	graph := Graph{
		Nodes: map[NodeID]*Node{},
		Edges: map[EdgeID]*Edge{},
	}
	return &graph
}

// GetAmountOfEdges returns the amount of edges in the graph
func (g *Graph) GetAmountOfEdges() (amount int) {
	return len(g.Edges)
}

// GetAmountOfNodes returns the amount of nodes in the graph
func (g *Graph) GetAmountOfNodes() (amount int) {
	return len(g.Nodes)
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(nodeID NodeID) (node *Node) {
	if g.CheckIfNodeExists(nodeID) {
		return g.GetNode(nodeID)
	} else {
		node := NewNode(nodeID)
		g.Nodes[node.Id] = node
	}
	return
}

// RemoveNode removes a node and all of its edges from the graph
func (g *Graph) RemoveNode(nodeID NodeID) {
	if g.CheckIfNodeExists(nodeID) {
		node := g.GetNode(nodeID)
		// remove incoming edges
		for _, edge := range node.Incoming {
			g.RemoveEdgeByID(edge.Id)
		}
		// remove outgoing edges
		for _, edge := range node.Outgoing {
			g.RemoveEdgeByID(edge.Id)
		}
		// delete node
		delete(g.Nodes, nodeID)
	}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to NodeID) *Graph {
	var fromNode, toNode *Node
	// from
	if !g.CheckIfNodeExists(from) {
		g.AddNode(from)
	}
	fromNode = g.GetNode(from)
	// to
	if !g.CheckIfNodeExists(to) {
		g.AddNode(to)
	}
	toNode = g.GetNode(to)
	// edge
	edgeID := GenerateEdgeID(fromNode, toNode)
	if !g.CheckIfEdgeExists(edgeID) {
		// create edge
		edge := NewEdge(fromNode, toNode, 0)
		g.Edges[edge.Id] = edge
	}
	return g
}

// RemoveEdge removes an edge from the graph
func (g *Graph) RemoveEdge(from, to NodeID) {
	// check if from node exists
	if !g.CheckIfNodeExists(from) {
		return
	}
	// to node exists
	if !g.CheckIfNodeExists(to) {
		return
	}
	// edge
	edgeID := generateEdgeIDFromNodeIDs(from, to)
	// remove edge
	g.RemoveEdgeByID(edgeID)
	return
}

// RemoveEdgeByID removes an edge from the graph using the edge id
func (g *Graph) RemoveEdgeByID(edgeID EdgeID) {
	if g.CheckIfEdgeExists(edgeID) {
		// remove edge
		delete(g.Edges, edgeID)
	}
}

// CheckIfNodeExists checks if a node exists in the graph
func (g *Graph) CheckIfNodeExists(nodeID NodeID) bool {
	_, ok := g.Nodes[nodeID]
	return ok
}

// CheckIfEdgeExists checks if an edge exists in the graph
func (g *Graph) CheckIfEdgeExists(edgeId EdgeID) bool {
	_, ok := g.Edges[edgeId]
	return ok
}

// CheckIfEdgeExistFromNodes checks if an edge exists in the graph
func (g *Graph) CheckIfEdgeExistFromNodes(fromNode, toNode NodeID) bool {
	from := g.GetNode(fromNode)
	if from == nil {
		return false
	}
	to := g.GetNode(toNode)
	if to == nil {
		return false
	}
	// edge
	edgeID := GenerateEdgeID(from, to)
	_, ok := g.Edges[edgeID]
	return ok
}

// GetNode returns a node from the graph
func (g *Graph) GetNode(nodeID NodeID) *Node {
	if g.CheckIfNodeExists(nodeID) {
		return g.Nodes[nodeID]
	} else {
		return nil
	}
}

// GetEdge returns an edge from the graph
func (g *Graph) GetEdge(edgeID EdgeID) *Edge {
	if g.CheckIfEdgeExists(edgeID) {
		return g.Edges[edgeID]
	} else {
		return nil
	}
}

// String returns a string representation of the graph
func (g *Graph) String() string {
	res := ""
	// iterate over nodes
	for _, n := range g.Nodes {
		res += "Node: " + string(n.Id) + "\n"
		// iterate over outgoing edges
		for _, e := range n.Outgoing {
			res += "\t" + string(n.Id) + "\t --> \t" + string(e.To.Id) + "\n"
		}
	}
	return res
}
