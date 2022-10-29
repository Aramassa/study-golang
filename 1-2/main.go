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
	Hoge
	v2 string
}

func (h *Hoge2) Bye() {
	fmt.Println("Bye " + h.v2 + ", Hello " + h.v1)
	h.v1 = "test"
}

func main() {
	hg1 := Hoge{v1: "Name01"}
	hg1.Hello()

	// あたかも継承のようなことはできるが、初期化方法がややこしい
	hg2 := Hoge2{Hoge{"Name02-001"}, "Name02"}
	hg2.Bye()

}
