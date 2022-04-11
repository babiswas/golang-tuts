package main
import "fmt"

func modify_value(x *int)int{
  *x=*x+1
  return *x
}

func main(){
  var x int
  x=3
  fmt.Println("The modified value is:",modify_value(&x))
}