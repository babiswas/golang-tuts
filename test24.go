package main
import "fmt"

type User struct{
  name string
  id int
}

func main(){
  user:=new(User)
  (*user).name="bapan"
  (*user).id=12
  fmt.Println(*user)
}