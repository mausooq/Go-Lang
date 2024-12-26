// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

func search(a [][]byte , word string) bool {
    row := len(a)
    col := len(a[0])
    
    var dfs func(r,c,idx int) bool
    dfs = func(r,c,idx int) bool{
        
        if idx == len(word){
            return true
        }
        
        if r < 0 ||  c < 0 || r >= row || c >= col {
            return false
        }
        
        if a[r][c] != word[idx] || a[r][c] == '#'{
            return false
        }
        
        temp :=  a[r][c]
        a[r][c] = '#'
        found := dfs(r-1,c,idx+1) || dfs(r,c+1,idx+1) || dfs(r+1,c,idx+1) || dfs(r,c-1,idx+1)
        
        a[r][c] = temp
        
        return found
        
    }
    
    for i :=0 ; i < row ; i++{
        for j := 0 ; j < col ; j++{
            if a[i][j] == word[0] && dfs(i,j,0){
                return true
            }
        }
    }
    
    return false 
}



func main() {
    a := [][]byte{
       {'A', 'B', 'C', 'E'},
	    {'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
    }
    word := "ABB"
    
    fmt.Println(search(a , word))
}