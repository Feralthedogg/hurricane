# Hurricane

Hurricane is a **3D spiral data structure** library implemented in Go.  
This project is **experimental** and is freely available for those who want to research and explore its concepts.

## 🚀 Overview
Hurricane provides a flexible 3D spiral data structure with customizable node connectivity strategies.  
It is designed with **goroutines** and **concurrency safety** in mind.  
This project is intended **for research and experimental purposes** and does not guarantee production stability.

## 🔥 Features
- Supports various spiral structures (Basic Spiral, Logarithmic Spiral, Archimedean Spiral)
- Customizable node adjacency strategies (simple linking, tree-like connections, center-based connections)
- **Concurrency-safe** using **RWMutex**
- Supports **BFS and DFS traversal algorithms**
- Layered node traversal for better data organization
- Spiral order sorting and traversal

## 🛠 Installation

```bash
go get github.com/Feralthedogg/hurricane
```

## 📚 Usage Example

```go
package main

import (
    "fmt"
    "math"
    "github.com/Feralthedogg/hurricane"
)

func main() {
    // Use default spiral function and adjacency strategy
    spiral := hurricane.DefaultSpiralFunc(0.2, math.Pi/10, 0.1)
    adjStrategy := hurricane.DefaultAdjacencyStrategy
    h := hurricane.NewHurricane("center", "center node", spiral, adjStrategy)

    // Add nodes
    for i := 1; i <= 10; i++ {
        nodeID := fmt.Sprintf("n%d", i)
        _, err := h.AddNode(nodeID, fmt.Sprintf("node %d", i))
        if err != nil {
            fmt.Println(err)
        }
    }

    // Spiral traversal
    fmt.Println("\nSpiral Traversal:")
    for _, node := range h.SpiralTraversal() {
        fmt.Printf("ID=%s, R=%.2f, Theta=%.2f, Z=%.2f\n", node.ID, node.R, node.Theta, node.Z)
    }

    // BFS traversal
    fmt.Println("\nBFS Traversal:")
    for _, node := range h.BFS(h.Center) {
        fmt.Printf("ID=%s\n", node.ID)
    }
}
```

## ⚡️ Running Benchmarks

To run the benchmark tests, use the following command:

```bash
go test -bench=.
```

## 📝 License

This project is licensed under the **MIT License**.  
Feel free to use it for research or experimental purposes.  
However, **the author is not responsible for any issues arising from its use**.

## 🤝 Contributing

This project is **open-source** and welcomes contributions for research and experimentation.  
Feel free to submit issues or pull requests to improve it.

---

> **Disclaimer:** This project is experimental and does not guarantee stability for production use.