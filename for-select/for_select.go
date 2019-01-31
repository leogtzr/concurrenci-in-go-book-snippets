package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	go func() {
		defer close(done)
		//done <- struct{}{}
		defer fmt.Println("Done goroutine :)")
		time.Sleep(5 * time.Second)
	}()

loop:
	for {
		select {
		case <-done: // Check if it is closed
			fmt.Println("Finally ... ")
			break loop
		default:
			fmt.Println("Nel ... ")
		}
		time.Sleep(1 * time.Second)
	}

}
