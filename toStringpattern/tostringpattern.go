package main

import (
	"fmt"
)

func take(n int, done <-chan interface{}, values ...int) <-chan int {
	valueStream := make(chan int)

	go func() {
		defer close(valueStream)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case valueStream <- values[i]:
			}
		}
	}()

	return valueStream
}

func main() {

	done := make(chan interface{})
	defer close(done)

	for v := range take(2, done, 54, 56, 23) {
		fmt.Println(v)
	}

}