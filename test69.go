package main
import "io"
import "fmt"
import "os"
import "bufio"

func main(){
   if len(os.Args)<2{
      fmt.Println("Provide path:")
      return
   }
   file,err:=os.Create(os.Args[1])
   if err!=nil{
      fmt.Println("Error occured")
      return
   }
   scanner:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter text")
   scanner.Scan()
   _,err=io.WriteString(file,scanner.Text())
   if err!=nil{
      fmt.Println("Error occured:")
      return
   }
   fmt.Println("operation completed")
}