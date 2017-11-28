package main

import "fmt"

func Person() func(int) int {
    sum := 0  // class property
    return func(x int) int {
        sum += x
        return sum
    }
}


func main(){
    person := Person()
    // instance of person
    fmt.Println(person(10)) // 10
    fmt.Println(person(12)) // 22
}
