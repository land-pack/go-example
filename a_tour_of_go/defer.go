package main

import "fmt"

func main() {

	defer fmt.Printf("Hello")
	fmt.Printf("World")

}
