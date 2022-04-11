package main
import "fmt"

func main(){
  arr:=[5]int{1,2,3,4,5}
  brr:=arr[1:4]
  fmt.Println(brr)
  x:=make([]int,5,10)
  for i:=0;i<len(x);i++{
     x[i]=10+i
  }
  for index,value:=range x{
    fmt.Println(index,value)
  }
  fmt.Println("The length of slice x is:",len(x))
  fmt.Println("The capacity of slice x is:",cap(x))
}