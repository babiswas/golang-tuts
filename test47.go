package main
import "fmt"
import "io"
import "log"
import "os"

func main(){
   r,w:=io.Pipe()

   go func(){
      fmt.Fprint(w,"Hello world")
      w.Close()
   }()

   if _,err:=io.Copy(os.Stdout,r);err!=nil{
      log.Fatal(err)
   }
}