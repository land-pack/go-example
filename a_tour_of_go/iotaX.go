package main

import "fmt"


type Car int

const (
    first Car = 1<< iota // 0
    second
    last
)


type ErrorCode int

const (

    _ = iota
    _
    _
    IllegalUser = iota

    PasswdError
    PrivalError
)



func main() {
    fmt.Println(second)
    fmt.Println(last)

    fmt.Println(IllegalUser)
    fmt.Println(PrivalError) // 2

}
