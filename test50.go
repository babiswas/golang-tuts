package main
import "fmt"
import "os"
import "io"

func check(err error){
  if err!=nil{
    panic(err)
  }
}

func main(){
   fmt.Println("Starting Operation:")
   file1,err:=os.Create("test13.txt")
   check(err)
   defer file1.Close()
   writer:=io.Writer(file1)
   writer.Write([]byte("hello World"))
   fmt.Println("Ending Operation")
}