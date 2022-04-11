package main
import "fmt"
import "sync"

func source1(ch1 chan<- string,wg *sync.WaitGroup){
   defer wg.Done()
   for{
      ch1<-"ping"
   }
}

func source2(ch2 chan<- string,wg *sync.WaitGroup){
   defer wg.Done()
   for{
     ch2<-"pong"
   }
}

func main(){
   var wg sync.WaitGroup
   ch1:=make(chan string)
   ch2:=make(chan string)
   go source1(ch1,&wg)
   go source2(ch2,&wg)
   for{
       select{
       case msg1:=<-ch1:
               fmt.Println(msg1)
       case msg2:=<-ch2:
               fmt.Println(msg2)
       default:
           fmt.Println("No message in pipeline")
       }
   }
}