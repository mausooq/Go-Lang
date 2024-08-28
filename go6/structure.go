package main
import ("fmt")

type Animal struct {
  name string
  age int
}

func main() {
  var cat Animal
//   var dog Animal

  cat.name = "haa"
  cat.age = 45
 


//   dog.name = "Cecilie"
//   dog.age = 24
 

//   fmt.Println("Name: ", cat.name)
//   fmt.Println("Age: ", cat.age)
 
//   fmt.Println("Name: ", dog.name)
//   fmt.Println("Age: ", dog.age)

  fmt.Print(cat)
}