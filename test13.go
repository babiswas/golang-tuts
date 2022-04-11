package main
import "fmt"

func sum_avg(arr [5]int)(int,float64){
   var sum int
   for _,value:=range arr{
     sum+=value
   }
   return sum,float64(sum)/float64(5)
}

func main(){
   arr:=[5]int{20,21,23,2,4}
   sum,avg:=sum_avg(arr)
   fmt.Println("The sum of the number is:",sum)
   fmt.Println("The average of the number is:",avg)
}