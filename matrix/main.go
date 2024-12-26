package main


import (
	"fmt"
)
func maxHistogramArea(height []int) int {
	stack := []int{}
	maxArea := 0
	index :=0
	for index < len(height){
		if len(stack) == 0 || height[index] >= height[stack[len(stack)-1]]{
			stack = append(stack,index)
			index++
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			
			w :=0
			if len(stack) ==0{
				w = index
			} else {
				w = index - stack[len(stack)-1]-1
			}
			maxArea = max(maxArea,height[top]*w)
		}
		
	}

	for len(stack) > 0{
		top := stack[len(stack)-1]

		stack = stack[:len(stack)-1]
		
		w :=0
		if len(stack) ==0{
			w = index
		} else {
			w = index - stack[len(stack)-1]-1
		}
		maxArea = max(maxArea,height[top]*w)
	}
	return maxArea
}

func maxRectangularArea(matrix [][]byte) int {
	
	row := len(matrix)
	col := len(matrix[0])
	 maxArea :=0
	 height := make([]int,col)
	for i :=0 ; i < row ; i++{
		for j := 0 ; j < col ; j++{
			if matrix[i][j] == 'Y'{
				height[j] += 1
			} else {
				height[j] = 0
				
			}
		}
		maxArea = max(maxArea,maxHistogramArea(height))
	}
return maxArea

}	
func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	matrix := [][]byte{
		{'Y', 'Y', 'N', 'Y'},
		{'Y', 'Y', 'N', 'Y'},
		{'Y', 'Y', 'Y', 'Y'},
		{'N', 'N', 'Y', 'Y'},
	}

	fmt.Println("Largest rectangular area of 'Y':", maxRectangularArea(matrix))
}


