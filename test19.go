package main
import "fmt"

func main(){
  add:=func(x,y int) int{
    return x+y
  }
  fmt.Println("The sum of 4 and 5 is:",add(4,5))
  fmt.Println("The sum of 10 and 4 is:",add(10,4))
}