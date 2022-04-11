package main
import "fmt"
import "time"

func test1(str1 string){
   for i:=0;i<10;i++{
     fmt.Println(str1)
     time.Sleep(time.Millisecond*2)
   }
}

func test2(str2 string){
   for i:=0;i<10;i++{
      fmt.Println(str2)
      time.Sleep(time.Millisecond*2)
   }
}

func main(){
   for i:=0;i<10;i++{
       go test1("hello1")
   }
   for i:=0;i<10;i++{
       go test2("bello")
   }
   var str1 string
   fmt.Scanln(&str1)
}