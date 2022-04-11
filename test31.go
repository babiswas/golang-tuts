package main
import "fmt"

func process_first(ch chan<- string){
   fmt.Println("Sending first message")
   ch<-"ping"
}

func process_second(ch chan<- string){
   fmt.Println("Sending second message")
   ch<-"pong"
}

func main(){
  var msg string
  c1:=make(chan string)
  c2:=make(chan string)
  go process_first(c1)
  go process_second(c2)
  for i:=0;i<2;i++{
    select{
      case msg=<-c1:
        fmt.Println("Processing channel1")
        fmt.Println(msg)
      case msg=<-c2:
        fmt.Println("Processing channel2")
        fmt.Println(msg)
    }
  }
}