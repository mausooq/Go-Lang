
package main
import "fmt"

func binary(a [][]int , k int) (int ,int){
    row := len(a)
    col := len(a[0])
    i:= 0 
    j := col -1
    for i < row && j >=0 {
        if a[i][j] == k {
            return i,j
        } else if  a[i][j] < k {
            i++
        } else {
            j--
        }
        
    }
    return -1,-1
}

func main() {
 a := [][]int{
     {1,2,3},
     {4,2,6},
     {1,8,9},
}

 row , col := binary(a,2)
 fmt.Println(row,col)
}
