// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

var dir =[][]int{
    {-1,0},
    {1,0},
    {0,-1},
    {0,1},
    {1,1},
    {-1,-1},
    {-1,1},
    {1,-1},
}

func shortestPathBinaryMatrix(grid [][]int) int{
    n := len(grid)
    
    if grid[0][0] != 0 || grid[n-1][n-1] != 0 {
        return -1
    }
    
    que := [][]int{{0,0,1}}
    grid[0][0] = 1 
    
    for len(que) > 0 {
        curr := que[0]
        que = que[1:]
        row,col,step := curr[0],curr[1],curr[2]
        
            
    	if row == n-1 && col == n-1 {
			return step
		}
		
		for _, val := range dir{
		    newRow , newCol := row+val[0] ,col+val[1]
		    
		    if newRow < n && newRow >= 0 && newCol >=0 && newCol < n && grid[newRow][newCol] == 0 {
		        que = append(que,[]int{newRow,newCol,step+1})
		        grid[newRow][newCol] = 1
		    }
		}
        
    }
    
    
    
    return -1
}


func main() {
 // Example 1
	grid1 := [][]int{
		{0, 1},
		{1, 0},
	}
	fmt.Println("Example 1:", shortestPathBinaryMatrix(grid1)) // Output: 2

	// Example 2
	grid2 := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}
	fmt.Println("Example 2:", shortestPathBinaryMatrix(grid2)) // Output: 4

	// Example 3
	grid3 := [][]int{
		{1, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}
	fmt.Println("Example 3:", shortestPathBinaryMatrix(grid3)) // Output: -1
}