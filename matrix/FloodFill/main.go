package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    row := len(image)
    col := len(image[0])
    k := image[sr][sc]

    // If the starting color is already the target color, no need to fill
    if k == color {
        return image
    }

    // Create a copy of the image to avoid modifying the original one
    a := make([][]int, row)
    for i := range image {
        a[i] = append([]int(nil), image[i]...)
    }

    var dfs func(r, c int)

    dfs = func(r, c int) {
        // Check if out of bounds or if the current pixel is not the original color
        if r < 0 || c < 0 || r >= row || c >= col || a[r][c] != k {
            return
        }

        // Set the current pixel to the new color
        a[r][c] = color

        // Recursively visit all 4 neighboring pixels
        dfs(r+1, c)
        dfs(r-1, c)
        dfs(r, c-1)
        dfs(r, c+1)
    }

    // Start DFS from the given starting point
    dfs(sr, sc)
    return a
}

func main() {
    // Example image
    image := [][]int{
        {1, 1, 1},
        {1, 1, 0},
        {1, 0, 1},
    }

    // Start point (sr, sc), new color
    sr, sc, color := 1, 1, 2

    // Call flood fill function
    result := floodFill(image, sr, sc, color)

    // Print the result
    fmt.Println("Filled Image:")
    for _, row := range result {
        fmt.Println(row)
    }
}

