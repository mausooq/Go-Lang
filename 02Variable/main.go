package main
import "fmt"



func main(){
	var username string ="harish"
	fmt.Println("username")
	fmt.Printf("variable is of type %T \n",username)

	var isLoggedIn bool = true
	fmt.Println("isLoggedIn")
	fmt.Printf("variable is of type %T \n",isLoggedIn)

	var smallval uint8 = 55
	fmt.Println("smallval")
	fmt.Printf("variable is of type %T \n",smallval)

	var smallfloat float32 = 55.99
	fmt.Println("smallfloat")
	fmt.Printf("variable is of type %T \n",smallfloat)

	//default value and alieass
	var num int
	fmt.Println("num")
	fmt.Printf("variable is of type %T \n",num)

	noOfuser := 3000
	fmt.Println(noOfuser)
}
