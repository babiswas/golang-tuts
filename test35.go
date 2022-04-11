package main
import "fmt"

func main(){
  jobs:=make(chan int,3)
  done:=make(chan bool)
  go func(){
     for{
        j,more:=<-jobs
        if more{
           fmt.Println("Received all jobs:",j)
        }else{
           fmt.Println("Received all jobs:")
           done<-true
           return
        }
     }
  }()
  for j:=1;j<=3;j++{
      jobs<-j
      fmt.Println("Sent jobs:",j)
  }
  close(jobs)
  fmt.Println("Sent all jobs:")
  fmt.Println(<-done) 
}