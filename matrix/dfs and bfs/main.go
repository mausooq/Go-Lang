package main

import "fmt"

func dfsTraversal(matrix [][]int, startRow, startCol int) {
	row := len(matrix)
	col := len(matrix[0])

	// DFS function to visit nodes and modify columns
	var dfs func(r, c int)
	dfs = func(r, c int) {
		// Boundary check and already visited check
		if r < 0 || c < 0 || r >= row || c >= col || matrix[r][c] == 0 {
			return
		}

		// Mark the entire column as 0
		for i := 0; i < row; i++ {
			matrix[i][c] = 0
		}

		// Explore all four directions
		dfs(r-1, c) // Up
		dfs(r+1, c) // Down
		dfs(r, c-1) // Left
		dfs(r, c+1) // Right
	}

	// Start DFS from the given starting point
	dfs(startRow, startCol)
}

func main() {
	// Example matrix
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Starting index
	startRow, startCol := 1, 1 // Start from the center of the matrix (value 5)

	fmt.Printf("Starting DFS from index (%d, %d)\n", startRow, startCol)
	dfsTraversal(matrix, startRow, startCol)

	// Print the modified matrix
	fmt.Println("Modified matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}
}
