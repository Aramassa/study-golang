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

type Hoge2 struct {
	v2 string
	Hoge
}

func (h *Hoge2) Bye() {
	fmt.Println("Bye " + h.v2)
	h.v1 = "test"
}

func main() {
	hg1 := Hoge2{v2: "Me2"}
	hg1.Hello()

}
