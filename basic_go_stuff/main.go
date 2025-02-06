package main 
import (
  "fmt"
  "flag"
)
type language string 
 func main()  {

  var lang string
  flag.StringVar(&lang,"language","eng","the required language")
  flag.Parse()

   greeting :=Greet(language(lang))
   fmt.Println(greeting)

 }

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
