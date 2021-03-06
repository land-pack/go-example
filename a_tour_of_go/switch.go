package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Runs on :")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("OS %s.", os)
	}
}
