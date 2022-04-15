package main
import "io"
import "log"
import "os"
import "fmt"

func check(err error){
   if err!=nil{
      log.Fatal(err)
      panic(err)
   }
}

func main(){
   file,err:=os.Open("test3.txt")
   defer file.Close()
   check(err)
   b,err:=io.ReadAll(file)
   check(err)
   fmt.Println(string(b))
}