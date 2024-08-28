package main
import ("fmt")

func main(){
	// a:= 5
	// b:= 10
	// fmt.Print(a+b)

	// var  (
	// 	sum1 = 200 + 2
	// 	sum2 = sum1 + 98
	// 	sum3 = sum1 + sum2
	// )
	// fmt.Println(sum3)

	// time := 22
	// if time < 10 {
	//   fmt.Println("Good morning.")
	// } else if time < 20 {
	//   fmt.Println("Good day.")
	// } else {
	//   fmt.Println("Good evening.")
	// }

// 	adj := [2]string{"big", "tasty"}
// 	fruits := [3]string{"apple", "orange", "banana"}
// 	for i:=0; i < len(adj); i++ {
// 	  for j:=0; j < len(fruits); j++ {
// 		fmt.Println(adj[i],fruits[j])
// }
// }

	// arr := [5]int{1,2,4,5,6}
	// for idx ,val := range arr{
	// 	fmt.Printf("%v\t%v\n", idx, val)
	// }
	arr := [5]int{1,2,4,5,6}
	for _ ,val := range arr{
		fmt.Printf("%v\n", val)
	}
}