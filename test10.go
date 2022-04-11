package main
import "fmt"

func main(){
  arr:=[]int{1,2,3,4,5}
  arr=append(arr,10,11,12)
  fmt.Println(arr)
  fmt.Println("The length and capacity of the slice is:")
  fmt.Println("The length of the slice is:",len(arr))
  fmt.Println("The capacity of the slice is:",cap(arr))
}