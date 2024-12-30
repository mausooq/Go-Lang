// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

var distance = [][]int{
    {-1, 0}, {1, 0}, {0, -1}, {0, 1},  
	{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
}

func FindWords(board [][]rune , str string) bool{
    row := len(board)
    col := len(board[0])
    
    
    var dfs func(r,c int ,idx int)bool
    dfs = func(r,c int ,idx int)bool{
    
        if idx == len(str){
            return true
        }
        if r < 0 || c < 0 || r >= row || c >= col || board[r][c] == '#' || board[r][c] != rune(str[idx]) {
            return false
        }
        
        
        box := board[r][c] 
        board[r][c] = '#'
        
        
        
        for _, d := range distance{
            newRow , newCol := r+d[0] , c+d[1]
                if dfs(newRow , newCol  , idx+1){
                    return true
                }
            
        }
        board[r][c] = box
        return false
    }
   
    for i :=0 ; i < row ;i++{
        for j:=0 ; j < col ; j++{
            if rune(str[0]) == board[i][j] && dfs(i,j,0){
                return true
            }
        }
    
    }
    
    return false
    
}




func main() {
	// Example board
	board := [][]rune{
		{'G', 'I', 'Z'},
		{'U', 'E', 'K'},
		{'Q', 'S', 'E'},
	}

	// Example dictionary
	dictionary := []string{"GEEKS", "FOR", "QUIZ", "GO"}

	// Find and print all valid words
	fmt.Println("Words found in the board:")
	 for _, str := range dictionary{
     if FindWords(board, str) {
            fmt.Println(str)
        }
	 }
	 
	
}