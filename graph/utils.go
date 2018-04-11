package graph

// Node struct to using in graphs
type Node struct {
	Value     interface{}
	Adjacency []*Node
}

// NewNode create a node
func NewNode(value interface{}) Node {
	return Node{value, []*Node{}}
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
