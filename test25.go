package main
import "fmt"

func add_func(add func(a,b int)int,c,d int)func()int{
    sum:=add(c,d)
    return func()int{
       return sum+5
    }
}

func main(){
   add:=func(a,b int)int{
      return a+b
   }
   fmt.Println("The sum of the numbers is:",add_func(add,4,5)())
}