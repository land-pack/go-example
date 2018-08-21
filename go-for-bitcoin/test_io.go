package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	s := strings.NewReader("hello Frank AK")
	r, err := ioutil.ReadAll(s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", r)
}
