package main

import "fmt"

func chanOwner() <-chan int {
	// since we know we'll produce six results, we can create a buffered channel
	// of five.
	resultStream := make(chan int, 5)
	go func() {
		defer close(resultStream)
		for i := 0; i <= 5; i++ {
			resultStream <- i
		}
	}()
	return resultStream
}

func main() {
	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
