package graph

// Node struct to using in graphs
type Node struct {
	Value     interface{}
	Adjacency []*Node
	edges     []*Edge
	pre       int
	pos       int
}

// Edge struct to using in graphs
type Edge struct {
	to     *Node
	weight int
}

// NewNode create a node
func NewNode(value interface{}) Node {
	return Node{value, []*Node{}, []*Edge{}, -99, -99}
}

// AddNeighbor include a neighbor to this node
func (n *Node) AddNeighbor(neighbor *Node, weight int) {
	n.Adjacency = append(n.Adjacency, neighbor)
	neighbor.Adjacency = append(neighbor.Adjacency, n)

	n.edges = append(n.edges, &Edge{neighbor, weight})
	neighbor.edges = append(neighbor.edges, &Edge{n, weight})
}

// AddNeighborDirected include a neighbor to this node
func (n *Node) AddNeighborDirected(neighbor *Node, weight int) {
	n.Adjacency = append(n.Adjacency, neighbor)
	n.edges = append(n.edges, &Edge{neighbor, weight})
}

// MarkAsVisited mark node as visited
func (n *Node) StartVisit(time int) {
	n.pre = time
}

// MarkAsVisited mark node as visited
func (n *Node) EndVisit(time int) {
	n.pos = time
}

// IsVisited get if the node was visited before
func (n *Node) IsVisited() bool {
	return n.pre != -99
}
