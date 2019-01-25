package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	go func() {
		messages <- "Ping"
	}()

	/*
		By default sends and receives block until both the sender and receiver are ready.
		This property allowed us to wait at the end of our program for the "ping" message
		without having to use any other synchronization.
	*/
	fmt.Println(<-messages)
}
