package main

import "fmt"

func minPathSum(grid [][]int) int {
    row := len(grid)
    col := len(grid[0])
    temp := make([][]int,row)
    for i := range temp {
        temp[i] = make([]int,col)
    }
    temp[0][0]= grid[0][0]
    for i :=1 ; i < col ; i++{
        temp[0][i] = temp[0][i-1] + grid[0][i]
    }
     for i :=1 ; i < row ; i++{
        temp[i][0] = temp[i-1][0] + grid[i][0]
    }
    for i := 1; i < row ; i++{
        for j := 1 ; j < col ; j++{
            temp[i][j] = min(temp[i-1][j],temp[i][j-1]) + grid[i][j]
        }
    }
    fmt.Println(temp)
    return temp[row-1][col-1]
}
func min(a , b int)int{
    if a < b {
        return a
    }
    return b
}

func main() {
    // Test case 1
    grid1 := [][]int{
        {1, 3, 1},
        {1, 5, 1},
        {4, 2, 1},
    }
    fmt.Println("Minimum Path Sum:", minPathSum(grid1)) // Output: 7

    // Test case 2
    grid2 := [][]int{
        {1, 2, 3},
        {4, 5, 6},
    }
    fmt.Println("Minimum Path Sum:", minPathSum(grid2)) // Output: 12
}
