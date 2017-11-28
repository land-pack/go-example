package main

import "fmt"
import "runtime"


func main() {
    go fmt.Println("Go")
    runtime.Gosched()



}
