package Hurricane

import (
	"fmt"
	"github.com/Feralthedogg/hurricane"
	"math"
)

func main() {
	spiral := hurricane.DefaultSpiralFunc(0.2, math.Pi/10, 0.1)
	adjStrategy := hurricane.DefaultAdjacencyStrategy
	h := hurricane.NewHurricane("center", "center node", spiral, adjStrategy)

	for i := 1; i <= 10; i++ {
		nodeID := fmt.Sprintf("n%d", i)
		_, err := h.AddNode(nodeID, fmt.Sprintf("node %d", i))
		if err != nil {
			fmt.Println(err)
		}
	}

	spiralOrder := h.SpiralTraversal()
	fmt.Println("Spiral Traversal:")
	for _, node := range spiralOrder {
		fmt.Printf("ID=%s, R=%.2f, Theta=%.2f, Z=%.2f\n", node.ID, node.R, node.Theta, node.Z)
	}

	layers := h.LayeredTraversal(0.2)
	fmt.Println("\nLayered Traversal:")
	for i, layer := range layers {
		fmt.Printf("Layer %d:\n", i)
		for _, node := range layer {
			fmt.Printf("  ID=%s, Z=%.2f\n", node.ID, node.Z)
		}
	}

	bfsOrder := h.BFS(h.Center)
	fmt.Println("\nBFS Traversal:")
	for _, node := range bfsOrder {
		fmt.Printf("ID=%s\n", node.ID)
	}

	dfsOrder := h.DFS(h.Center)
	fmt.Println("\nDFS Traversal:")
	for _, node := range dfsOrder {
		fmt.Printf("ID=%s\n", node.ID)
	}
}
