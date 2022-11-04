package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	messages := make(chan string, 2)

	messages <- "ping1"
	messages <- "ping2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	done := make(chan bool, 1)
	go worker(done)

	<-done // これをコメントアウトするとどうなるか確認してみよう。
}
