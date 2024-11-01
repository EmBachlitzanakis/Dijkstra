package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Edge represents a directed edge with a destination node and weight
type Edge struct {
	to     int
	weight int
}

// Graph is represented as an adjacency list
type Graph map[int][]Edge

// Item represents an item in the priority queue
type Item struct {
	node     int // node index
	priority int // priority (distance from start)
	index    int // index of the item in the heap
}

// PriorityQueue implements a priority queue using min-heap
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority of an item in the queue
func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

// Dijkstra function finds the shortest paths from a start node
func Dijkstra(graph Graph, start int) map[int]int {
	// Initialize distance map with infinity for each node
	dist := make(map[int]int)
	for node := range graph {
		dist[node] = math.MaxInt64
	}
	dist[start] = 0

	// Priority queue for tracking minimum distance
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{node: start, priority: 0})

	for pq.Len() > 0 {
		// Get node with smallest distance
		item := heap.Pop(&pq).(*Item)
		u := item.node
		distance := item.priority

		// Skip if we found a better path already
		if distance > dist[u] {
			continue
		}

		// Relax edges
		for _, edge := range graph[u] {
			v := edge.to
			weight := edge.weight
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(&pq, &Item{node: v, priority: dist[v]})
			}
		}
	}

	return dist
}

// Helper function to add edges to the graph
func addEdge(graph Graph, from, to, weight int) {
	graph[from] = append(graph[from], Edge{to: to, weight: weight})
}

func main() {
	// Create a graph
	graph := make(Graph)
	addEdge(graph, 1, 2, 2)
	addEdge(graph, 1, 3, 4)
	addEdge(graph, 2, 3, 1)
	addEdge(graph, 2, 4, 7)
	addEdge(graph, 3, 5, 3)
	addEdge(graph, 4, 6, 1)
	addEdge(graph, 5, 4, 2)
	addEdge(graph, 5, 6, 5)

	// Find shortest paths from start node 1
	dist := Dijkstra(graph, 1)

	// Print shortest distances
	for node, d := range dist {
		fmt.Printf("Distance from 1 to %d: %d\n", node, d)
	}
}
