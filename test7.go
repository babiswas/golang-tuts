package main
import "fmt"
import "os"

func main(){
   if len(os.Args)<2{
     fmt.Println("Please provide command line arguments:")
     return
   }
   for index,value:=range os.Args[1:]{
      fmt.Println(index,value)
   }
}