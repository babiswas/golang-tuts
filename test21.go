package main
import "fmt"

func sum_arr(arr [6]int)int{
   defer fmt.Println("Addition of numbers completed")
   sum:=0
   for _,value:=range arr{
      sum+=value
   }
   return sum
}

func main(){
   defer fmt.Println("Main completed the task")
   arr:=[6]int{12,14,15,10,21,22}
   sum:=sum_arr(arr)
   fmt.Println("The sum of the numbers is:",sum)
}