package main
import "fmt"
import "sync"
import "strconv"

func send_message(ch chan<- string,wg *sync.WaitGroup){
     defer wg.Done()
     for i:=0;i<10;i++{
        ch<-"hello"+strconv.Itoa(i)
     }
     close(ch)
}

func main(){
   var msg string
   var wg sync.WaitGroup
   ch:=make(chan string)
   wg.Add(1)
   go send_message(ch,&wg)
   for i:=0;i<10;i++{
      msg=<-ch
      fmt.Println(msg)
   }
}