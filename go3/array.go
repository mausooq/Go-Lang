package main
import ("fmt")


func main() {
	// var arr1 = [4]int{1, 2, 3, 4}
	// arr2 := [3]string{"ppp","pppppp","ppppppppp"}

	// var arr1 = [...]int{1, 2, 3, 4}
	// arr2 := [...]string{"ppp","pppppp","ppppppppp"}

	
	// fmt.Println(arr1)
	// fmt.Println(arr2)


//   prices := [3]int{10,20,30}

//   fmt.Println(prices[0])
//   fmt.Println(prices[2])

	// prices[2] = 50
  	// fmt.Println(prices)

  	// fmt.Println(len(prices))

	//   myslice1 := []int{1, 2, 3, 4, 5, 6}
	//   fmt.Printf("myslice1 = %v\n", myslice1)
	//   fmt.Printf("length = %d\n", len(myslice1))
	//   fmt.Printf("capacity = %d\n", cap(myslice1))
	
	//   myslice1 = append(myslice1, 20, 21)
	//   fmt.Printf("myslice1 = %v\n", myslice1)
	//   fmt.Printf("length = %d\n", len(myslice1))
	//   fmt.Printf("capacity = %d\n", cap(myslice1))

// 	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
//   myslice1 := arr1[1:5] // Slice array
//   fmt.Printf("myslice1 = %v\n", myslice1)
//   fmt.Printf("length = %d\n", len(myslice1))
//   fmt.Printf("capacity = %d\n", cap(myslice1))

//   myslice1 = arr1[1:3] // Change length by re-slicing the array
//   fmt.Printf("myslice1 = %v\n", myslice1)
//   fmt.Printf("length = %d\n", len(myslice1))
//   fmt.Printf("capacity = %d\n", cap(myslice1))

//   myslice1 = append(myslice1, 20, 21, 22, 23,24,15,44,55,44) // Change length by appending items
//   fmt.Printf("myslice1 = %v\n", myslice1)
//   fmt.Printf("length = %d\n", len(myslice1))
//   fmt.Printf("capacity = %d\n", cap(myslice1))


numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
fmt.Println(numbers)

needNum := numbers[:len(numbers)-10]
numcopy := make([]int,len(needNum))
copy(numcopy,needNum)
fmt.Print(numcopy)
}

