package main

import (
	"2-1/internal/example"
	pkg "2-1/pkg/example"
	"fmt"
)

func main() {
	fmt.Println("Hello")

	ex := example.Example{Name: "Yosuke Ex of intenrl", Age: 41}
	fmt.Println(ex.MyName())

	ex2 := pkg.Example{Name: "Yosuke Ex of pkg", Age: 41}
	fmt.Println(ex2.MyName())
}
