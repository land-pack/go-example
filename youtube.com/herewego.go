package main

import "fmt"


// Comment
/* Start here */


func main() {
    fmt.Println("Hello World!")

    var age int64 = 40
    var favNum float64 = 1.622422
    randNum := 1
    fmt.Println(randNum)

    //randNum = "hello"

    fmt.Println(age, " ", favNum)
    fmt.Printf("The type is %T\n", randNum)


    i := 1
    for i <= 10 {
        fmt.Println(i)
        i++ // i -
    }

    for j :=0 ; j < 5 ; j ++ {
        fmt.Printf("The j is %d \n" ,j)
    }


    yourAge := 19

    switch yourAge {
        case 16: fmt.Println("you should study at school !")
        case 18: fmt.Println("you can driver car !")
        case 20: fmt.Println("you can do anything")

        default: fmt.Println("Default age")
    }


    var favNums2[5] float64
    favNums2[1] = 90
    favNums2[2] = 190
    for i , v := range favNums2 {
        fmt.Println(i, v)

    }

    // slice a array
    numSlice := [] int { 5,3,2,2,1,1}

    numSlice2 := numSlice[3:5]

    fmt.Println("numSlice2[2] =", numSlice2[1])
}
