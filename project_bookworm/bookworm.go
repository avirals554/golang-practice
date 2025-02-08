package main
import "fmt"
import "os"
import "path/filepath"
func main(){
  f,err:=os.open(home/aviral/learning_go/project_bookworm/testdata/bookworm.json)
  if err!=nil{
    fmt.Println("unsucessful",err)

  }
 defer f.close()



}
