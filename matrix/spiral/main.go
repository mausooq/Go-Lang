package main

import "fmt"

// Function to traverse a matrix in spiral order
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	// Initialize boundaries
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	result := []int{}

	for top <= bottom && left <= right {
		// Traverse from left to right along the top row
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// Traverse from top to bottom along the right column
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// Traverse from right to left along the bottom row (if not already traversed)
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		// Traverse from bottom to top along the left column (if not already traversed)
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

func main() {
	// Example matrix
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	// Print the original matrix
	fmt.Println("Original Matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	// Get the spiral order
	spiral := spiralOrder(matrix)

	// Print the spiral order
	fmt.Println("Spiral Order:")
	fmt.Println(spiral)
}

