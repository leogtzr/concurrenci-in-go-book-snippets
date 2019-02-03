package main

import "fmt"

func take(num int, done chan interface{}, values ...int) <-chan interface{} {
	intStream := make(chan interface{})

	go func() {
		defer close(intStream)

		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case intStream <- values[i]:
			}
		}

	}()

	return intStream
}

func main() {
	done := make(chan interface{})
	defer close(done)

	for v := range take(2, done, 1, 4, 5, 8) {
		fmt.Println(v)
	}
}
