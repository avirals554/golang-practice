package main

import "fmt"
func main(){
	greeting := Greet("eng")
	fmt.Println(greeting)
}
type language string
func Greet( l language) string{
 switch l {
 case "eng":
  return "hello"
 case"fr":
  return"Bonjour le monde"
default:
  return""
   
 }
}



