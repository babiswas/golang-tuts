package main
import "encoding/json"
import "fmt"

type User struct{
  Name string `json:"name"`
  Id int `json:"id"`
}

func main(){
  user:=User{"bapan",22}
  data,err:=json.Marshal(user)
  if err!=nil{
    fmt.Println("Error occured",err)
    return
  }
  fmt.Println(string(data))
  data1:=json.RawMessage(string(data))
  if err!=nil{
      fmt.Println(err)
      return
  }
  user1:=new(User)
  bytedata,err: =data1.MarshalJSON()
  err=json.Unmarshal(bytedata,&user1)
  fmt.Println(*user1)
}