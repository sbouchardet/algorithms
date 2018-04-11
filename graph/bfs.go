package graph

import "fmt"

//BFS breadth first search
func BFS(graph []*Node, value interface{}) {
	// n := len(graph)
	S := []*Node{graph[0]}
	for len(S) > 0 {
		nodeToVisitedNow := S[0]
		if nodeToVisitedNow.Value == value {
			fmt.Print("Finded!")
		}
		nodeToVisitedNow.MarkAsVisited()
		S = S[1:]
		adjs := nodeToVisitedNow.Adjacency
		for i := 0; i < len(adjs); i++ {
			if !adjs[i].IsVisited() {
				S = append(S, adjs[i])
			}
		}
	}
}
