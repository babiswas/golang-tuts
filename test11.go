package main
import "fmt"

func main(){
  mynums:=[]float64{1.2,3.45,5.6,6.5,4.2,3.88,9.25}
  var sum float64
  sum=0.0
  for _,value:=range mynums{
     sum+=value
  }
  fmt.Println("The sum of the numbers is:",sum)
  fmt.Println("The average of the numbers is:",sum/float64(len(mynums)))
}