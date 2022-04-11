package main
import "fmt"

func main(){
   var arr [5]int
   for i:=0;i<5;i++{
     arr[i]=i
   }
   fmt.Println(arr)
   fmt.Println("Traditional looping")
   for i:=0;i<len(arr);i++{
     fmt.Println(arr[i])
   }
   fmt.Println("Golang style looping")
   for index,val:=range(arr){
      fmt.Println(index,val)
   }
   brr:=[5]int{1,2,3,4,5}
   fmt.Println("Displaying array in golang style looping")
   for index,value:=range brr{
      fmt.Println(index,value)
   }
}