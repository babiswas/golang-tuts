package main
import "fmt"
import "os"

func main(){
   if len(os.Args)<2{
     fmt.Println("Plese provide path:")
     return
   }
   data,err:=os.ReadFile(os.Args[1])
   if err!=nil{
      fmt.Println("Error occured")
      return
   }
   fmt.Println(string(data))
   err=os.WriteFile(os.Args[1],[]byte("hello world helloworld"),0666)
   if err!=nil{
     fmt.Println("Error occured")
     return
   }
   data,_=os.ReadFile(os.Args[1])
   fmt.Println(string(data))
}