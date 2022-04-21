package main
import "fmt"
import "os"
import "bufio"

func main(){
   if len(os.Args)<2{
      fmt.Println("Please provide path")
      return
   }
   data,err:=os.ReadFile(os.Args[1])
   if err!=nil{
      fmt.Println("Error occured")
      return
   }
   fmt.Println(string(data))
   scanner:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter the line")
   scanner.Scan()
   err=os.WriteFile(os.Args[1],[]byte(scanner.Text()),0666)
   if err!=nil{
      fmt.Println("Error Occured")
      return
   }
   data,err=os.ReadFile(os.Args[1])
   if err!=nil{
      fmt.Println("Error occured")
      return
   }
   fmt.Println(string(data))
}