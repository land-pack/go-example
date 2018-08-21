package main


import "fmt"

type ErrorCodeType int16

const (
    succ ErrorCodeType = iota
    MissSomething ErrorCodeType = iota
    DieServer
    _
    _
    defaultCode
)

func main() {

    fmt.Println(defaultCode)
}
