package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Pair structure for the priority queue
type Pair struct {
	node, distance int
}

// Priority Queue implementation
type PriorityQueue []Pair

func (pq PriorityQueue) Len() int { 
	return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Pair))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// Dijkstra's algorithm function
func dijkstra(graph [][]int, src int) []int {
	n := len(graph)
	distances := make([]int, n)
	for i := range distances {
		distances[i] = math.MaxInt32 
	}
	distances[src] = 0

	// Priority Queue
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Pair{src, 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(Pair)
		node := current.node
		distance := current.distance

		// If the current distance is already greater, skip
		if distance > distances[node] {
			continue
		}

		// Process neighbors
		for neighbor := 0; neighbor < n; neighbor++ {
			if graph[node][neighbor] > 0 { // There's an edge
				newDist := distances[node] + graph[node][neighbor]
				if newDist < distances[neighbor] {
					distances[neighbor] = newDist
					heap.Push(pq, Pair{neighbor, newDist})
				}
			}
		}
	}

	return distances
}

func main() {
	// Adjacency Matrix of a Weighted Graph
	graph := [][]int{
		{0, 2, 4, 0, 0},
		{2, 0, 1, 7, 0},
		{4, 1, 0, 3, 5},
		{0, 7, 3, 0, 1},
		{0, 0, 5, 1, 0},
	}

	source := 0
	distances := dijkstra(graph, source)

	fmt.Printf("Shortest distances from node %d:\n", source)
	for i, d := range distances {
		fmt.Printf("Node %d: %d\n", i, d)
	}
}
