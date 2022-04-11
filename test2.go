package main
import "fmt"
import "sync"

func ping_generator(ch chan<-string,wg *sync.WaitGroup){
   defer wg.Done()
   for{
     ch <-"ping"
   }
}

func pong_generator(ch chan<- string,wg *sync.WaitGroup){
   defer wg.Done()
   for{
     ch <-"pong"
   }
}

func read_ping_pong(ch <-chan string,wg *sync.WaitGroup){
   defer wg.Done()
   mydata:=""
   for{
     mydata=<-ch
     fmt.Println(mydata)
   }
}

func main(){
   var wg sync.WaitGroup
   ch:=make(chan string)
   wg.Add(1)
   go ping_generator(ch,&wg)
   wg.Add(1)
   go pong_generator(ch,&wg)
   wg.Add(1)
   go read_ping_pong(ch,&wg)
   wg.Wait()
   
}