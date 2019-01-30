package main

import (
	"fmt"
	"time"
)

func main() {
	var c <-chan int
	select {
	case <-c:
		fmt.Println("Done :D")
	case <-time.After(5 * time.Second):
		fmt.Println("We have timed out ... ")
	}

	fmt.Println("done ... :) bye")
}
