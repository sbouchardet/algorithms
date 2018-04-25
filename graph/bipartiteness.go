package graph

func BipartitenessTest(g *Node) bool {
	bfsTree := BFS(g)

	treeGroupedByLevel := groupByLeve(bfsTree)

	for i := 0; i < len(treeGroupedByLevel); i++ {
		sameLevel := treeGroupedByLevel[i]
		if edgeInVector(sameLevel) {
			return false
		}
	}

	return true

}

func edgeInVector(vec []*TreeNode) bool {
	for i := 0; i < len(vec); i++ {
		thisNodeAdj := vec[i].GraphNode.Adjacency
		for j := i + 1; j < len(vec); j++ {
			for k := 0; k < len(thisNodeAdj); k++ {
				adj := thisNodeAdj[k]
				if adj == vec[j].GraphNode {
					return true
				}
			}
		}
	}

	return false
}

func groupByLeve(allNodes []*TreeNode) [][]*TreeNode {
	maxLevel := getMaxLevel(allNodes)
	result := [][]*TreeNode{{}}

	for i := 0; i <= maxLevel; i++ {
		result = append(result, []*TreeNode{})
	}
	for j := 0; j < len(allNodes); j++ {
		node := allNodes[j]
		result[node.Level] = append(result[node.Level], node)
	}

	return result

}

func getMaxLevel(allNodes []*TreeNode) int {
	max := 0
	for i := 0; i < len(allNodes); i++ {
		if max < allNodes[i].Level {
			max = allNodes[i].Level
		}
	}
	return max
}
