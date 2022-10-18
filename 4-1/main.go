package main

import (
	"fmt"
	"strconv"
	"time"
)

type Gr struct {
	Name string
}

func (g *Gr) MyName() string {
	return g.Name
}

func (g *Gr) PrintMyName() {
	fmt.Println(g.Name)
}

func main() {
	var g1 = func(msg string) {
		for s := 1; s <= 3; s++ {
			fmt.Printf("Hello World : %s - %s \n", msg, strconv.Itoa(s))
		}
	}
	g := Gr{Name: "Yosuke"}

	for i := 1; i <= 10; i++ {
		go g1(strconv.Itoa(i))
		go g.PrintMyName()
	}

	time.Sleep(time.Second)

}
