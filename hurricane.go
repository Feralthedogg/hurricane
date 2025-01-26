// hurricane.go

package Hurricane

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

type Node struct {
	ID       string  // Node ID
	Data     string  // Node data
	R        float64 // Radius (spiral structure)
	Theta    float64 // Rotation angle (radians)
	Z        float64 // Height (Z-axis)
	X        float64 // X coordinate
	Y        float64 // Y coordinate
	Adjacent []*Node // Adjacent nodes
}

type Hurricane struct {
	sync.RWMutex
	Center      *Node             // Center node
	Nodes       []*Node           // List of all nodes
	NodeMap     map[string]*Node  // ID to Node mapping
	spiralFunc  SpiralFunc        // Spiral calculation function
	adjStrategy AdjacencyStrategy // Adjacency connection strategy
}

type SpiralFunc func(index int) (r, theta, z float64)

type AdjacencyStrategy func(h *Hurricane, newNode *Node)

func NewHurricane(centerID, centerData string, spiral SpiralFunc, adjStrategy AdjacencyStrategy) *Hurricane {
	centerNode := &Node{
		ID:    centerID,
		Data:  centerData,
		R:     0.0,
		Theta: 0.0,
		Z:     0.0,
		X:     0.0,
		Y:     0.0,
	}

	h := &Hurricane{
		Center:      centerNode,
		Nodes:       []*Node{centerNode},
		NodeMap:     map[string]*Node{centerID: centerNode},
		spiralFunc:  spiral,
		adjStrategy: adjStrategy,
	}
	return h
}

func (h *Hurricane) AddNode(id, data string) (*Node, error) {
	h.Lock()
	defer h.Unlock()

	if _, exists := h.NodeMap[id]; exists {
		return nil, fmt.Errorf("node %q already exists", id)
	}

	nodeIndex := len(h.Nodes)
	r, theta, z := h.spiralFunc(nodeIndex)
	x := r * math.Cos(theta)
	y := r * math.Sin(theta)

	newNode := &Node{
		ID:    id,
		Data:  data,
		R:     r,
		Theta: theta,
		Z:     z,
		X:     x,
		Y:     y,
	}

	h.Nodes = append(h.Nodes, newNode)
	h.NodeMap[id] = newNode

	if h.adjStrategy != nil {
		h.adjStrategy(h, newNode)
	}

	return newNode, nil
}

func (h *Hurricane) GetNode(id string) *Node {
	h.RLock()
	defer h.RUnlock()
	return h.NodeMap[id]
}

func (h *Hurricane) SpiralTraversal() []*Node {
	h.RLock()
	defer h.RUnlock()

	sortedNodes := make([]*Node, len(h.Nodes))
	copy(sortedNodes, h.Nodes)

	sort.Slice(sortedNodes, func(i, j int) bool {
		if sortedNodes[i].R == sortedNodes[j].R {
			return sortedNodes[i].Theta < sortedNodes[j].Theta
		}
		return sortedNodes[i].R < sortedNodes[j].R
	})
	return sortedNodes
}

func (h *Hurricane) LayeredTraversal(layerSize float64) [][]*Node {
	h.RLock()
	defer h.RUnlock()

	layerMap := make(map[int][]*Node)
	maxLayer := 0

	for _, node := range h.Nodes {
		layer := int(math.Floor(node.Z / layerSize))
		layerMap[layer] = append(layerMap[layer], node)
		if layer > maxLayer {
			maxLayer = layer
		}
	}

	layers := make([][]*Node, maxLayer+1)
	for i := 0; i <= maxLayer; i++ {
		layers[i] = layerMap[i]
	}
	return layers
}

func (h *Hurricane) BFS(start *Node) []*Node {
	h.RLock()
	defer h.RUnlock()

	visited := make(map[string]bool)
	queue := []*Node{start}
	visited[start.ID] = true
	result := []*Node{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, adj := range node.Adjacent {
			if !visited[adj.ID] {
				visited[adj.ID] = true
				queue = append(queue, adj)
			}
		}
	}
	return result
}

func (h *Hurricane) DFS(start *Node) []*Node {
	h.RLock()
	defer h.RUnlock()

	visited := make(map[string]bool)
	stack := []*Node{start}
	result := []*Node{}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[node.ID] {
			visited[node.ID] = true
			result = append(result, node)

			for i := len(node.Adjacent) - 1; i >= 0; i-- {
				adj := node.Adjacent[i]
				if !visited[adj.ID] {
					stack = append(stack, adj)
				}
			}
		}
	}
	return result
}

func (h *Hurricane) GetAllNodes() []*Node {
	h.RLock()
	defer h.RUnlock()

	nodesCopy := make([]*Node, len(h.Nodes))
	copy(nodesCopy, h.Nodes)
	return nodesCopy
}
