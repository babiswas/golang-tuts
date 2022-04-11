package main
import "fmt"
import "sync"

func test1(ch1 chan<- string,wg *sync.WaitGroup){
   defer wg.Done()
   ch1<-"pong"
}

func test2(ch1 <-chan string,ch2 chan<- string,wg *sync.WaitGroup){
  var msg string
  defer wg.Done()
  msg=<-ch1
  ch2<-msg
}

func main(){
  var wg sync.WaitGroup
  ch1:=make(chan string,1)
  ch2:=make(chan string,1)
  wg.Add(1)
  go test1(ch1,&wg)
  wg.Add(1)
  go test2(ch1,ch2,&wg)
  fmt.Println(<-ch2)
  wg.Wait()
}