package main

import "fmt"
import "time"
import "errors"

// MyError is an error implementation that includes a time and message

type MyError struct {
    When time.Time
    What string
}


func (e MyError) Error() string {

    return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
    return MyError{
        time.Date(2018,3,15,22,30,0,0, time.UTC),
        "the file system has gone away!",
   }
}

func main(){
    if err := oops(); err != nil {
        fmt.Println(err)
    }
    // system errors
    err := errors.New("emit macho dwarf: elf header corrupted\n")
    if err != nil {
        fmt.Print(err)
    }
}
