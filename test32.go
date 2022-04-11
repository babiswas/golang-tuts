package main
import "fmt"
import "time"
import "sync"

func worker1(ch chan<-string,wg *sync.WaitGroup){
  defer wg.Done()
  time.Sleep(time.Second)
  ch<-"pong"
}

func main(){
   var wg sync.WaitGroup
   ch1:=make(chan string,1)
   wg.Add(1)
   go worker1(ch1,&wg)
   select{
      case msg:=<-ch1:
         fmt.Println(msg)
      case <-time.After(1 * time.Millisecond):
         fmt.Println("Timeout occured")
   }
   wg.Wait()
}