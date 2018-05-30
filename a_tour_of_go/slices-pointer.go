package main

import "fmt"

func main() {

	names := [4]string{
		"hello",
		"world",
		"is",
		"beginer",
	}

	b := names[2:4]
	b[1] = "XXX"

	fmt.Println(b)
	fmt.Println(names)
}
