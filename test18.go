package main
import "fmt"

type User struct{
  name string
  id int
}

type Record struct{
  name string
  rowno int
}

func process_nything(obj ...interface{}){
   for _,val:=range obj{
      fmt.Println(val)
   }
}

func main(){
   user:=User{"bapan",21}
   record:=Record{"databse",32}
   record1:=Record{"database1",33}
   user1:=User{"bapan20",22}
   process_nything(user,record,record1,user1)
}