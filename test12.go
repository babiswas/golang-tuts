package main
import "fmt"

func add(arr [5]int)int{
  sum:=0
  for _,value:=range arr{
    sum+=value
  }
  return sum
}

func main(){
   arr:=[5]int{10,11,12,13,14}
   sum:=add(arr)
   fmt.Println("The sum of the numbers in the array is:",sum)
   fmt.Println("The average of the numbers is:",sum/len(arr))
}