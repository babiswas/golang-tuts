package main
import "fmt"
import "sync"
import "time"

func worker1(ch chan<- string,wg *sync.WaitGroup){
  defer wg.Done()
  time.After(3*time.Millisecond)
  ch<-"ping"
}

func worker2(ch chan<-string,wg *sync.WaitGroup){
   defer wg.Done()
   time.After(3*time.Millisecond)
   ch<-"pong"
}

func main(){
  var wg sync.WaitGroup
  var msg string
  ch:=make(chan string,1)
  wg.Add(1)
  go worker1(ch,&wg)
  select{
      case msg=<-ch:
        fmt.Println(msg)
      case <-time.After(1*time.Millisecond)
        fmt.Println("Timeout happened")
  }
  wg.Add(1)
  go worker2(ch,&wg)
  select{
      case msg=<ch:
         fmt.Println(msg)
      case<-time.After(10*time.Millisecond):
         fmt.Println("Timeout happend")
  }
  wg.Wait()
}