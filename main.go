package main 
import "fmt"
 func main()  {
  greeting :=Greet("eng")
   fmt.Println(greeting)

 }
type language string
var phrasebook =map[language]string{
 "eng" : "hello world",
 "fr"  : "bonjour le monde",
 "hin" : "hum kisi ",

}
func Greet(l language) string{
  greeting,ok:=phrasebook[l ]
  if !ok{
    return fmt.Sprintf("unsuported language %q" ,l)
  }
return greeting
}
