package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{91, 22}
	p := &v

	fmt.Println(p.X)

}
