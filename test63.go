package main
import "os"
import "fmt"
import "io"

func main(){
   if len(os.Args)<2{
     fmt.Println("Please provide path")
     return
   }
   file,err:=os.Open(os.Args[1])
   if err!=nil{
     fmt.Println("Error occured:")
     return
   }
   defer file.Close()
   file1,err:=os.Create(os.Args[2])
   if err!=nil{
      fmt.Println("Error occured")
      return
   }
   defer file1.Close()
   io.Copy(file1,file)
   fmt.Println("Operation completed")
}