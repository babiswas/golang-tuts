package main
import "fmt"

func sum_number(add func(x,y int)int,x,y int)func()int{
    z:=5
    return func()int{
       z=z+1
       return add(x,y)+z
    }
}

func main(){
   add1:=func(x,y int)int{
      return x+y
   }
   add2:=sum_number(add1,3,4)
   fmt.Println("The sum of the number is:",add2())
   fmt.Println("The sum of the number is:",add2())
   fmt.Println("The sum of the number is:",add2())
   fmt.Println("The sum of the number is:",add2())
}