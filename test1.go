package main
import "fmt"

func send_numbers(ch chan<- int){
   for i:=0;i<10;i++{
     ch<-i
   }
   close(ch)
}

func recive_numbers(ch <-chan int){
   num:=0
   for i:=0;i<10;i++{
      num=<-ch
      fmt.Println(num)
   }
}

func main(){
  i:=0
  ch:=make(chan int)
  go send_numbers(ch)
  go recive_numbers(ch)
  for{
    i=i+1
    if i==1000000000{
       break
    }
  }
}