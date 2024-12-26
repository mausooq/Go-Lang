package main

import "fmt"

func dfs(graph [][]int, start int) {
	n := len(graph)
	visited := make([]bool, n)
	stack := []int{start}

	visited[start] = true

	for len(stack) > 0 {
		// Dequeue the first node
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", node)

		// Enqueue all unvisited neighbors
		for neighbor, isEdge := range graph[node] {
			if isEdge == 1 && !visited[neighbor] {
				stack = append(stack, neighbor)
				visited[neighbor] = true
			}
		}
	}
}

func bfsMatrix(graph [][]int, start int) {
	n := len(graph)
	visited := make([]bool, n)
	queue := []int{start}

	visited[start] = true

	for len(queue) > 0 {
		// Dequeue the first node
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node)

		// Enqueue all unvisited neighbors
		for neighbor, isEdge := range graph[node] {
			if isEdge == 1 && !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}
}

func main() {
	// Adjacency Matrix
	graph := [][]int{
		{0, 1, 1, 0, 0},
		{1, 0, 1, 1, 0},
		{1, 1, 0, 0, 1},
		{0, 1, 0, 0, 1},
		{0, 0, 1, 1, 0},
	}

	fmt.Println("BFS Traversal:")
	bfsMatrix(graph, 0) // Start BFS from node 0
	dfs(graph,0)
	
}
