package main
import "fmt"
import "log"
import "os/exec"

func main(){
    pycmd:=exec.Command("py")
    opres,err:=pycmd.Output()
    if err!=nil{
       log.Fatal(err)
       return
    }
    fmt.Println(string(opres))
}