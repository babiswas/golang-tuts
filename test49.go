package main
import "fmt"
import "io"
import "os"
import "log"

func check(err error){
  if err!=nil{
     log.Fatal(err)
     panic(err)
  }
}

func main(){
   fmt.Println("Started the operation")
   file1,err:=os.Open("test1.txt")
   check(err)
   defer file1.Close()
   file2,err:=os.Create("test10.txt")
   check(err)
   defer file2.Close()
   io.Copy(file2,file1)
}