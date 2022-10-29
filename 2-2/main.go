package main

import (
	"fmt"

	"github.com/Aramassa/study-golang/2-1/pkg/example"
)

func main() {
	ex := example.Example{MyName: "Test"}
	fmt.Println(ex.MyName())
}
