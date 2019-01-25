package main

import (
	"fmt"
	"sync"
)

func main() {
	var msg string
	var count int
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan string)

	ping := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < 10; i++ {
			count++
			c <- fmt.Sprintf("ping-%d", count)
			fmt.Println(<-c)
		}
	}

	pong := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < 10; i++ {
			msg = <-c
			fmt.Println(msg)
			count++
			c <- fmt.Sprintf("pong-%d", count)
		}
	}

	wg.Add(2)
	close(begin)
	go ping()
	go pong()

	wg.Wait()
	fmt.Println("Bye ... ")

}
