package main

import "fmt"
import "sync"


func main() {
    var wg sync.WaitGroup
    wg.Add(3)

    go func (){
        fmt.Println("Go!")
        wg.Done()
    }()

    go func (){
        fmt.Println("Go!")
        wg.Done()
    }()

    go func (){
        fmt.Println("Go!")
        wg.Done()
    }()

    wg.Wait()
}
