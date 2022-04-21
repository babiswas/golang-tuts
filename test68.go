package main
import "io"
import "fmt"
import "os"

func main(){
   if len(os.Args)>2{
      fmt.Println("Please provide a path")
      return
   }
   file,err:=os.OpenFile(os.Args[1],os.O_RDWR|os.O_CREATE,0666)
   if err!=nil{
     fmt.Println("Error occured")
     return
   }
   data,err:=io.ReadAll(file)
   if err!=nil{
     fmt.Println(err)
     return
   }
   fmt.Println(string(data))
}