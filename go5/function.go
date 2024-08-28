// package main
// import ("fmt")

// func main(){
// 	myMessage()
// 	myMessage()
// 	myMessage()
// 	myMessage()
// }

// func 	myMessage(){
// 	fmt.Println("Hello, World!")
// }
package main
import ("fmt")

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
  a, b := myFunction(5, "Hello")
  fmt.Println(a)
  fmt.Println(b)
}

// package main
// import ("fmt")

// func factorial(x int)(y int){
// 	if x == 0 {
// 		y =1
// 	} else{
// 		y = x * factorial(x-1)
// 	}
// 	return
// }
// func main(){
// 	fmt.Println(factorial(5))
// }