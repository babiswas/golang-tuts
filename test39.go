package main
import "time"
import "fmt"
import "sync"

func worker(id int,jobs <-chan int,results chan<- int,wg *sync.WaitGroup){
  defer wg.Done()
  for j:=range jobs{
     fmt.Println("Worker",id,"Started job",j)
     time.Sleep(time.Second)
     fmt.Println("Worker",id,"Finished job",j)
     results<-j*2 
  }
}

func main(){
   var wg sync.WaitGroup
   const numJobs=5
   jobs:=make(chan int,numJobs)
   result:=make(chan int,numJobs)
   for i:=1;i<=5;i++{
      wg.Add(1)
      go worker(i,jobs,result,&wg)
   }
   for j:=1;j<=numJobs;j++{
      jobs<-j
   }
   close(jobs)
   for a:=1;a<=5;a++{
     fmt.Println(<-result)
   }
   wg.Wait()
}