package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	p = &Vertex{12, 22}
)

func main() {

	v := Vertex{X: 2} // Y default 0
	w := Vertex{}     // Both default as 0

	fmt.Println(v, w, *p)

}
