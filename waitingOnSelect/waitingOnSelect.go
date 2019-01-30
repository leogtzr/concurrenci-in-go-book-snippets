package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	go func() {
		defer close(done)
		time.Sleep(5 * time.Second)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
			// Simulate work
			workCounter++
		}
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)

}
