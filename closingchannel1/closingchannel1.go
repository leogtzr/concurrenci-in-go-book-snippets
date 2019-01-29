package main

import "fmt"

func main() {
	c := make(chan int)
	// Closing channel:
	close(c)

	if _, ok := <-c; !ok {
		fmt.Println("Channel is closed")
	}

}
