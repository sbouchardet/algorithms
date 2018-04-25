package graph

type TreeNode struct {
	GraphNode *Node
	Sons      []*TreeNode
	Level     int
}

//BFS breadth first search
func BFS(root *Node) []*TreeNode {
	rootTreeNode := TreeNode{root, []*TreeNode{}, 0}
	allTreeNodes := []*TreeNode{&rootTreeNode}
	S := []*TreeNode{&rootTreeNode}
	time := 0
	for len(S) > 0 {
		nodeToVisitedNow := S[0]
		graphNodeToVisit := nodeToVisitedNow.GraphNode
		graphNodeToVisit.StartVisit(time)
		time = time + 1
		S = S[1:]
		adjs := graphNodeToVisit.Adjacency
		for i := 0; i < len(adjs); i++ {
			if !adjs[i].IsVisited() {
				adjTreeNode := &TreeNode{adjs[i], []*TreeNode{}, nodeToVisitedNow.Level + 1}
				nodeToVisitedNow.Sons = append(nodeToVisitedNow.Sons, adjTreeNode)
				allTreeNodes = append(allTreeNodes, adjTreeNode)
				S = append(S, adjTreeNode)
			}
		}
	}

	return allTreeNodes

}
