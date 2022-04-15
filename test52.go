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
  file1,err:=os.Open("test1.txt")
  check(err)
  defer file1.Close()
  file2,err:=os.Create("test11.txt")
  check(err)
  defer file2.Close()
  buf:=make([]byte,300)
  n,err:=io.CopyBuffer(file2,file1,buf)
  check(err)
  fmt.Printf("Copied %d",n)
}