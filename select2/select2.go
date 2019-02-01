package main

import "fmt"

func main() {
	var x int = 2
	c := make(chan interface{})
	close(c)

	for i := 0; i < 10; i++ {
		select {
		case <-c:
			fmt.Println("Available ... ")
			x++
		}
	}

	fmt.Println(x)
}
