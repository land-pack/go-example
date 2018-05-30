package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}

	a := s[1:4]

	b := s[2:]
	c := s[0:3]
	// d := s[4:2]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	// fmt.Println(d)

}
