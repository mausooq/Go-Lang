package main

import (
	"fmt"
)

func maximalSquare(matrix [][]byte) int {
	// Handle edge case for an empty matrix
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	maxSquare := 0
	row := len(matrix)
	col := len(matrix[0])

	// Initialize DP table
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}

	// Traverse the matrix
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					// First row or column, just copy the value
					dp[i][j] = 1
				} else {
					// Calculate minimum of the three neighbors
					dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
					// fmt.Println(dp[i][j])
				}
				// Update maxSquare
				maxSquare = max(maxSquare, dp[i][j])
			}
		}
	}

	// Return the area of the largest square
	return maxSquare * maxSquare
}

// Helper function to find the minimum of three integers
func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example matrix
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	// Call the function and print the result
	fmt.Println("Maximal Square Area:", maximalSquare(matrix))
}
