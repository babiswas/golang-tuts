package main
import "fmt"
import "sync"

func main(){
   var wg sync.WaitGroup
   ch1:=make(chan interface{})
   wg.Add(1)
   go func(ch chan<- interface{}){
     defer wg.Done()
     obj:=[]interface{}{1,"hello","bello",3.4}
     for i:=0;i<len(obj);i++{
        ch1<-obj[i]
     }
     close(ch1)
   }(ch1)
   for msg:=range ch1{
      fmt.Println(msg)
   }
   wg.Wait()
}