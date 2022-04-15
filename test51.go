package main
import "io"
import "fmt"
import "os"

func check(err error){
  if err!=nil{
    panic(err)
  }
}

func main(){
   file1,err:=os.Open("test1.txt")
   check(err)
   defer file1.Close()
   n,err:=io.CopyN(os.Stdout,file1,15)
   check(err)
   fmt.Printf("\nThe number of bytes read is %d:",n)
}