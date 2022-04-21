package main
import "fmt"
import "os"
import "io"

func main(){
  if len(os.Args)<2{
    fmt.Println("Please provide path")
    return
  }
  file,err:=os.OpenFile(os.Args[1],os.O_RDWR|os.O_CREATE,0666)
  if err!=nil{
    fmt.Println("Error occured")
    return
  }
  defer file.Close()
  io.Copy(os.Stdout,file)
  fmt.Printf("\nOperation Completed")
  err=os.WriteFile(os.Args[1],[]byte("hello world"),0666)
  if err!=nil{
    fmt.Println("Error occured:")
    return
  }
  io.Copy(os.Stdout,file)
}