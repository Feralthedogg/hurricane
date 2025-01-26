// hurricane_test.go

package Hurricane

import (
	"fmt"
	"math"
	"testing"
)

func BenchmarkAddNode_DefaultSpiral_DefaultAdjacency(b *testing.B) {
	spiral := DefaultSpiralFunc(0.2, math.Pi/10, 0.1)
	hurricane := NewHurricane("center", "center node", spiral, DefaultAdjacencyStrategy)

	for i := 0; i < b.N; i++ {
		nodeID := fmt.Sprintf("n%d", i)
		_, err := hurricane.AddNode(nodeID, "data")
		if err != nil {
			b.Fatalf("Failed to add node: %v", err)
		}
	}
}

func BenchmarkAddNode_LogSpiral_ConnectToCenter(b *testing.B) {
	spiral := LogSpiralFunc(0.1, 0.2, 0.05)
	hurricane := NewHurricane("center", "center node", spiral, ConnectToCenterStrategy)

	for i := 0; i < b.N; i++ {
		nodeID := fmt.Sprintf("n%d", i)
		_, err := hurricane.AddNode(nodeID, "data")
		if err != nil {
			b.Fatalf("Failed to add node: %v", err)
		}
	}
}

func BenchmarkSpiralTraversal(b *testing.B) {
	spiral := DefaultSpiralFunc(0.2, math.Pi/10, 0.1)
	hurricane := NewHurricane("center", "center node", spiral, DefaultAdjacencyStrategy)

	for i := 1; i <= 1000; i++ {
		nodeID := fmt.Sprintf("n%d", i)
		_, err := hurricane.AddNode(nodeID, "data")
		if err != nil {
			b.Fatalf("Failed to add node: %v", err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hurricane.SpiralTraversal()
	}
}

func BenchmarkLayeredTraversal(b *testing.B) {
	spiral := DefaultSpiralFunc(0.2, math.Pi/10, 0.1)
	hurricane := NewHurricane("center", "center node", spiral, DefaultAdjacencyStrategy)

	for i := 1; i <= 1000; i++ {
		nodeID := fmt.Sprintf("n%d", i)
		_, err := hurricane.AddNode(nodeID, "data")
		if err != nil {
			b.Fatalf("Failed to add node: %v", err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hurricane.LayeredTraversal(0.2)
	}
}

func BenchmarkBFS(b *testing.B) {
	spiral := DefaultSpiralFunc(0.2, math.Pi/10, 0.1)
	hurricane := NewHurricane("center", "center node", spiral, DefaultAdjacencyStrategy)

	for i := 1; i <= 1000; i++ {
		nodeID := fmt.Sprintf("n%d", i)
		_, err := hurricane.AddNode(nodeID, "data")
		if err != nil {
			b.Fatalf("Failed to add node: %v", err)
		}
	}

	startNode := hurricane.GetNode("center")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hurricane.BFS(startNode)
	}
}

func BenchmarkDFS(b *testing.B) {
	spiral := DefaultSpiralFunc(0.2, math.Pi/10, 0.1)
	hurricane := NewHurricane("center", "center node", spiral, DefaultAdjacencyStrategy)

	for i := 1; i <= 1000; i++ {
		nodeID := fmt.Sprintf("n%d", i)
		_, err := hurricane.AddNode(nodeID, "data")
		if err != nil {
			b.Fatalf("Failed to add node: %v", err)
		}
	}

	startNode := hurricane.GetNode("center")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hurricane.DFS(startNode)
	}
}
