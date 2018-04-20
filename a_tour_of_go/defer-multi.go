package main

import "fmt"

func main() {

	defer fmt.Printf("Counting ..")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)

	}

	defer fmt.Println("Defer done")
}
