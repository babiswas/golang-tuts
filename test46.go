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
    file,err:=os.Create("test3.txt")
    defer file.Close()
    check(err)
    var line string
    line="Hello world"
    n,err:=io.WriteString(file,line)
    check(err)
    fmt.Printf("Written %d bytes",n)
}