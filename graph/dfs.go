package graph

// DFS Deaph First Search
func DFS(root *Node, value interface{}, startTime int) int{
	root.StartVisit(startTime)
	var time = startTime + 1
	for i:=0;i<len(root.Adjacency);i++ {
		adj := root.Adjacency[i]
		if (!adj.IsVisited()) {
			time = DFS(root.Adjacency[i], value, time)
		}
	}
	root.EndVisit(time)

	return time + 1

}
