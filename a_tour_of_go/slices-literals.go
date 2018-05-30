package main

import "fmt"

func main() {

	q := []int{1, 2, 2, 2, 3, 4, 45, 6}
	fmt.Println(q)

	b := []bool{true, false, false, true}
	fmt.Println(b)

	s := []struct {
		x int
		b bool
	}{
		{2, false},
		{19, true},
		{1, false},
	}

	fmt.Println(s)

}
