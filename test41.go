package main
import "time"
import "fmt"

func main(){
   req:=make(chan int,5)
   for i:=1;i<5;i++{
       req<-i
   }
   close(req)
   limiter:=time.Tick(20000*time.Millisecond)
   for ind:=range req{
     <-limiter
     fmt.Println("Request",ind,time.Now())
   }
   burstyLimiter:=make(chan time.Time,3)
   for i:=0;i<3;i++{
     burstyLimiter<-time.Now()
   }
   go func(){
     for t:=range time.Tick(20000*time.Millisecond){
         burstyLimiter<-t
     }
   }()
   burstyRequest:=make(chan int,5)
   for i:=1;i<=5;i++{
     burstyRequest<-i
   }
   close(burstyRequest)
   for req:=range burstyRequest{
     <-burstyLimiter
     fmt.Println("request",req,time.Now())
   }
}