package main

import "fmt"

func sum(s []int, c chan int) {
    _sum := 0
    for _, v := range s {
        _sum += v
    }
    c <- _sum // send sum to c
}

func main() {
    s := []int{7, 5, 36, 22, 1, 8}
    c := make(chan int)

    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)

    x, y := <-c, <-c // receive from c
    fmt.Println(x, y, x+y)
}
