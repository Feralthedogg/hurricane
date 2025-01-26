// adjacency.go

package Hurricane

func DefaultAdjacencyStrategy(h *Hurricane, newNode *Node) {
	if len(h.Nodes) > 1 {
		prevNode := h.Nodes[len(h.Nodes)-2]
		connectNodes(prevNode, newNode)
	}
}

func ConnectToCenterStrategy(h *Hurricane, newNode *Node) {
	if h.Center != newNode {
		connectNodes(h.Center, newNode)
	}
}

func TreeAdjacencyStrategy(h *Hurricane, newNode *Node) {
	if len(h.Nodes) > 1 {
		prevNode := h.Nodes[len(h.Nodes)-2]
		connectNodes(h.Center, newNode)
		connectNodes(prevNode, newNode)
	} else {
		connectNodes(h.Center, newNode)
	}
}

func connectNodes(a, b *Node) {
	a.Adjacent = append(a.Adjacent, b)
	b.Adjacent = append(b.Adjacent, a)
}
