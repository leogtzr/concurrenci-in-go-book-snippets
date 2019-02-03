package main

import "fmt"

func functor(done chan interface{}, fn func(int) int, values ...int) <-chan int {
	intStream := make(chan int)

	go func() {
		defer close(intStream)

		for _, v := range values {
			select {
			case <-done:
				return
			case intStream <- fn(v):
			}
		}

	}()

	return intStream
}

func main() {
	done := make(chan interface{})
	defer close(done)

	for n := range functor(
		done,
		func(x int) int {
			return x + 1
		}, 3, 5, 2,
	) {
		fmt.Println(n)
	}

}
