package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {

	v := Vertex{12, 22}
	v.X = 4
	fmt.Println(v)

}
