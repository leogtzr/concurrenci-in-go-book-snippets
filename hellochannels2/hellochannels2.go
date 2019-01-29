package main

import (
	"fmt"
	"time"
)

func hello(stringStream chan<- string) {
	time.Sleep(1 * time.Second)
	stringStream <- "Hello!"
}

func main() {
	stringStream := make(chan string)
	go hello(stringStream)
	fmt.Println(<-stringStream)
}
