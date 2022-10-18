package main

import (
	"fmt"
)

type Hoge struct {
	v1 string
}

func (h *Hoge) Hello() {
	fmt.Println("Hello " + h.v1)
}

func main() {
	var hg1 = Hoge{v1: "Me"}
	hg1.Hello()

	var hg2 = Hoge{v1: "You"}
	hg2.Hello()

	hg1.Hello()
}
