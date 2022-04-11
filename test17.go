package main
import "fmt"

func add(arr ...int)int{
  sum:=0
  for _,value:=range arr{
    sum+=value
  }
  return sum
}

func main(){
   sum:=add(1,2,3,4,5,6,7,8)
   fmt.Println("The sum of the numbers returned by variadic function is:",sum)
}