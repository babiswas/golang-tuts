package main
import "fmt"

func test1()(a int){
  var c int
  var b int 
  c=2
  b=2
  a=c+b+10
  return
}

func main(){
  fmt.Println(test1())
}