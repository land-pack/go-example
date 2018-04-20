package main

import (
	"fmt"
)

func main() {
	var f float64= 19.22
	i := int64(f)

	fmt.Printf("The origin type %T | value=%v || after conversion type=%T | value=%v\n", f, f, i, i)
}
