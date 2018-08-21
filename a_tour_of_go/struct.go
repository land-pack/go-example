package main

import "fmt"

type V struct {
	Y int `hasbdsak`
	X int `sdfnlsaf`
	S string `lsflsam`
}

func main() {

	fmt.Println(V{12, 22, "hello"})
    //Struct fileds are accessed using a dot.
    fmt.Println("X is ", V{98, 22, "a"}.X)

}
