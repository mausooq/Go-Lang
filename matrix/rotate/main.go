package main

import "fmt"

// Function to rotate a matrix 90 degrees clockwise
func rotateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	if n == 0 || len(matrix[0]) == 0 {
		return nil // Handle empty matrix
	}

	// Step 1: Create a new matrix for the result
	rotated := make([][]int, n)
	for i := range rotated {
		rotated[i] = make([]int, n)
	}

	// Step 2: Rotate the matrix 90 degrees clockwise
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated[j][n-i-1] = matrix[i][j]
		}
	}

	return rotated
}

func main() {
	// Example matrix
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Print the original matrix
	fmt.Println("Original Matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	// Rotate the matrix
	rotatedMatrix := rotateMatrix(matrix)

	// Print the rotated matrix
	fmt.Println("Rotated Matrix (90° Clockwise):")
	for _, row := range rotatedMatrix {
		fmt.Println(row)
	}
}
