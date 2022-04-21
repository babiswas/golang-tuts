package main
import "fmt"
import "os"
import "io"

func main(){
   if len(os.Args)>2{
      fmt.Println("Provide path:")
      return
   }
   file,err:=os.Create(os.Args[1])
   if err!=nil{
      fmt.Println("Error occured:")
      return
   }
   defer file.Close()
   writer:=io.Writer(file)
   writer.Write([]byte("hello world"))
}