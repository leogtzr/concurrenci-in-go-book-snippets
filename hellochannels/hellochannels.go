package main

import (
	"fmt"
	"time"
)

func main() {
	stringStream := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		stringStream <- "Hello channels!"
	}()

	// This blocks until the channel has data ...
	fmt.Println(<-stringStream)

}
