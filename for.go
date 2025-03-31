package main

import "fmt"


func main(){
    
    for i:=1;i<6;i++{
        fmt.Println("this is me",i,"times")
    }
    for i := range 3 {
        fmt.Println(i)
    }
    for{
        fmt.Println("kkko")
        break
    }
}
