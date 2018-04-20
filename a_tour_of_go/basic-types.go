package main

import "fmt"
import "math/cmplx"

var (
	ToBe   bool       = false
	MaxInt int64      = 1<<32 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	fmt.Printf("The type is %T value = %v\n", ToBe, ToBe)
	fmt.Printf("The type is %T value = %v\n", MaxInt, MaxInt)
	fmt.Printf("The type is %T value = %v\n", z, z)
}
