package main

import "fmt"

func generator(done chan interface{}, ints ...int) <-chan int {
	intStream := make(chan int)

	go func() {
		defer close(intStream)
		for _, i := range ints {
			select {
			case <-done:
				return
			case intStream <- i:
			}
		}
	}()

	return intStream
}

func multiplier(done chan interface{}, ints <-chan int, mult int) <-chan int {
	intStream := make(chan int)

	go func() {
		defer close(intStream)
		for i := range ints {
			select {
			case <-done:
				return
			case intStream <- i * mult:
			}
		}
	}()

	return intStream
}

func adder(done chan interface{}, ints <-chan int, add int) <-chan int {
	intStream := make(chan int)

	go func() {
		defer close(intStream)
		for i := range ints {
			select {
			case <-done:
				return
			case intStream <- i + add:
			}
		}
	}()

	return intStream
}

func main() {
	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 4, 3, 5, 2)
	pipeline := multiplier(done, adder(done, multiplier(done, intStream, 2), 1), 2)

	for n := range pipeline {
		fmt.Println(n)
	}
}
