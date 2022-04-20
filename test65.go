package main
import "fmt"
import "os"

func main(){
  if len(os.Args)<2{
    fmt.Println("Please provide file path")
    return
  }
  file,err:=os.Create(os.Args[1])
  if err!=nil{
    fmt.Println("Error occured:")
    return
  }
  b:=[]byte("Hello world")
  err=os.WriteFile(os.Args[1],b,0666)
  if err!=nil{
    fmt.Println("Error occured:")
    return
  }
  file.Close()
  data,err:=os.ReadFile(os.Args[1])
  if err!=nil{
    fmt.Println("Error occured")
    return
  }
  fmt.Println(string(data))
  fmt.Println("Operation Completed:")
}