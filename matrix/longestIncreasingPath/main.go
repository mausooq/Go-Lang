// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

var direction = [][]int{
     {0,1},
     {0,-1},
     {1,0},
     {-1,0},
 }

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
     row := len(matrix)
     col := len(matrix[0])
     
    temp := make([][]int , row)
    for i := range temp {
        temp[i] = make([]int,col)
    }
    
    var dfs func(r,c int) int 
    dfs = func (r,c int) int{
        
        if temp[r][c] != 0 {
            return temp[r][c]
        }
        
        maxPath :=  1
        
        for _, val := range direction{
            newRow , newCol := r+val[0] , c+val[1]
            if newRow >= 0 && newCol >=0 && newRow < row && newCol < col && matrix[r][c] < matrix[newRow][newCol]{
                maxPath = max(maxPath,1+dfs(newRow,newCol))
            }
        }
        temp[r][c] = maxPath
        // fmt.Println(temp)
        return maxPath

    }
    
    
    maxPath := 0
    for i := 0 ; i < row ; i++{
        for j :=0 ; j < col ; j++{
            maxPath = max(maxPath,dfs(i,j))
        }
    }
    return maxPath
}
func max(a , b int ) int {
    if a > b {
        return a
    }
    return b
}
 

func main() {
	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}

	fmt.Println("Longest Increasing Path:", longestIncreasingPath(matrix)) // Output: 4
}