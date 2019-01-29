package main

import (
	"fmt"
	"time"
)

func main() {

	stringStream := make(chan string)

	go func() {
		fmt.Println("Hello I am #1")
		fmt.Println(<-stringStream, "#1")
	}()

	go func() {
		fmt.Println("Hello I am #2")
		fmt.Println(<-stringStream, "#2")
	}()

	go func() {
		fmt.Println("Hello I am #3")
		fmt.Println(<-stringStream, "#3")
	}()

	time.Sleep(2 * time.Second)
	stringStream <- "Que putos"

	time.Sleep(2 * time.Second)

}
