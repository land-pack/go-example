package main


import (
    "fmt"
    "os/user"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic("No user found")
    }

    fmt.Println(user)

    // Get user home path
    fmt.Println(user.HomeDir)

}
