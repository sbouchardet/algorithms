package graph

import (
	"fmt"
)

type piFunc struct {
	n  *Node
	pi int
}

// FindBestPath Using Dijskstr algorithms to find bestPath
func FindBestPath(root *Node) []*Edge {
	fmt.Print("HI")
	//TODO: implement algorithm.
	return nil
}

func getBestEdge(piVector []*piFunc) (*piFunc, int, []*piFunc) {
	var min *piFunc
	iMin := -1
	for i, value := range piVector {
		if min != nil && value.pi < min.pi {
			min = value
			iMin = i
		}
	}
	piVectorRemoved := append(piVector[:iMin], piVector[iMin+1:]...)
	return min, iMin, piVectorRemoved
}
