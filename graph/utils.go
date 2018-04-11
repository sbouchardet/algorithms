package graph

// Node struct to using in graphs
type Node struct {
	Value     interface{}
	Adjacency []*Node
	visited   bool
}

// NewNode create a node
func NewNode(value interface{}) Node {
	return Node{value, []*Node{}, false}
}

// AddNeighbor include a neighbor to this node
func (n *Node) AddNeighbor(neighbor *Node) {
	n.Adjacency = append(n.Adjacency, neighbor)
	neighbor.Adjacency = append(neighbor.Adjacency, n)
}

// AddNeighborDirected include a neighbor to this node
func (n *Node) AddNeighborDirected(neighbor *Node) {
	n.Adjacency = append(n.Adjacency, neighbor)
}

// MarkAsVisited mark node as visited
func (n *Node) MarkAsVisited() {
	n.visited = true
}

// IsVisited get if the node was visited before
func (n *Node) IsVisited() bool {
	return n.visited
}
