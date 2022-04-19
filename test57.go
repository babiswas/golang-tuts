package main
import "fmt"
import "strconv"
import "sync"
import "time"

func test1(wg *sync.WaitGroup,ch1 chan<- string){
  defer wg.Done()
  for i:=0;i<10;i++{
     ch1<-"Hello "+strconv.Itoa(i)
     time.Sleep(8*time.Millisecond)
  }
  close(ch1)
}

func test2(wg *sync.WaitGroup,ch2 chan<- string){
  defer wg.Done()
  for i:=0;i<10;i++{
      ch2<-"Bello "+strconv.Itoa(i)
      time.Sleep(10*time.Millisecond)
  }
  close(ch2)
}

func main(){
  var msg string
  var ok1 bool
  var ok2 bool
  ch1:=make(chan string,5)
  ch2:=make(chan string,5)
  var wg sync.WaitGroup
  fmt.Println("Starting operation:")
  wg.Add(2)
  go test1(&wg,ch1)
  go test2(&wg,ch2)
  for{
     select{
         case msg,ok1=<-ch1:
           fmt.Println(msg)
           if ok1{
             for m:=range ch1{
                fmt.Println(m)
             }
           }
         case msg,ok2=<-ch2:
           if ok2{
              for m:=range ch2{
                fmt.Println(m)
             }
           }
     }
     if !ok1 && !ok2{
        break
     }
  }
  wg.Wait()
}