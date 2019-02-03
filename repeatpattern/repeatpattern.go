package main

import "fmt"

func repeat(done chan interface{}, values ...interface{}) <-chan interface{} {
	intStream := make(chan interface{})

	go func() {
		defer close(intStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case intStream <- v:
				}
			}
		}
	}()

	return intStream
}

func main() {
	done := make(chan interface{})
	defer close(done)

	for v := range repeat(done, 1, 4, 5) {
		fmt.Println(v)
	}
}
