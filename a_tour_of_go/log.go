package main 


import "log"
import "fmt"


func main(){
    log.Fatal("It's hit me")
    // Because the Fatal will call `os.Exit` after output log
    // so you can't call any function at the rest of code
    fmt.Println("You can't see this message")

    // What if you call Panic
    //log.Panic("It's a panic ")
    //fmt.Println("You can't see this message")

}
