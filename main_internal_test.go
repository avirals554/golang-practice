package main
import "testing"
func TestGreet(t *testig.T){
  want:="hello"
  got := Greet()
// This is test comment
  if got !=want{
    t.Errorf("expected:%q,got:%q",want,got)
  }
        

}


