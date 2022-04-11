package main
import "fmt"


func display(arr ...int){
  for index,value:=range arr{
      fmt.Println(index,value)
  }
}

func display_arr(arr []int){
  for index,value:=range arr{
      fmt.Println(index,value)
  }
}

func display_str(brr ...string){
   for index,value:=range brr{
      fmt.Println(index,value)
   }
}

func main(){
  arr:=[]int{1,2,3,4,5}
  fmt.Println("Displaying variable number of integers in array:")
  display(1,2,3,4,5,6,7,8,9,10)
  fmt.Println("Passing slice in function:")
  display_arr(arr)
  fmt.Println("Displaying variable number of strings:")
  display_str("hello","bello","mello","tello","nello","dello")
}