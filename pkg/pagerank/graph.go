package pagerank

type Graph struct {
	Nodes map[NodeID]*Node
	Edges map[EdgeID]*Edge
}

func NewGraph() (g *Graph) {
	graph := Graph{
		Nodes: map[NodeID]*Node{},
		Edges: map[EdgeID]*Edge{},
	}
	return &graph
}

func (g *Graph) AddNode(nodeID string) (node *Node) {
	if g.CheckIfNodeExists(nodeID) {
		return g.GetNode(nodeID)
	} else {
		node := NewNode(nodeID)
		g.Nodes[node.Id] = node
	}
	return
}

func (g *Graph) AddEdge(from, to string) *Graph {
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
	if !g.CheckIfEdgeExists(string(edgeID)) {
		// create edge
		edge := NewEdge(fromNode, toNode, 0)
		g.Edges[edge.Id] = edge
	}
	return g
}

func (g *Graph) RemoveEdge(from, to string) {
	// TODO
}

func (g *Graph) CheckIfNodeExists(nodeID string) bool {
	_, ok := g.Nodes[NodeID(nodeID)]
	return ok
}

func (g *Graph) CheckIfEdgeExists(edgeId string) bool {
	_, ok := g.Edges[EdgeID(edgeId)]
	return ok
}

func (g *Graph) GetNode(nodeID string) *Node {
	if g.CheckIfNodeExists(nodeID) {
		return g.Nodes[NodeID(nodeID)]
	} else {
		return nil
	}
}

func (g *Graph) GetEdge(edgeID string) *Edge {
	if g.CheckIfEdgeExists(edgeID) {
		return g.Edges[EdgeID(edgeID)]
	} else {
		return nil
	}
}

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
