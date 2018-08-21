package main


import "fmt"


func main(){
    ch1 := make(chan int, 1)
    ch2 := make(chan int, 2)


    ch1 <- 1
    ch2 <- 2
    for {
         select  {
             case <-ch1:
                 fmt.Println("Trigger channel 1")
             case <-ch2:
                 fmt.Println("Trigger channel 2")
 //            default:
 //                fmt.Println("Default event")
         }
    }
}
