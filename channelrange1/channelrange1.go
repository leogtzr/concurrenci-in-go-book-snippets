package main

import (
	"fmt"
	"time"
)

func main() {
	intStream := make(chan int)

	go func() {
		defer close(intStream) // This is a common pattern:
		for i := 1; i <= 5; i++ {
			if i == 3 {
				time.Sleep(3 * time.Second)
			}
			intStream <- i
		}
	}()

	for integer := range intStream { // iterates until the channel is closed
		fmt.Printf("%v ", integer)
	}

	fmt.Println()
}
