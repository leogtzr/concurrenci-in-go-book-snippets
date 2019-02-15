package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork(done <-chan interface{},
	id int,
	wg *sync.WaitGroup,
	result chan<- int,
) {
	started := time.Now()
	defer wg.Done()

	// Simulate random load:
	simulatedLoadTime := time.Duration(1+rand.Intn(5)) * time.Second
	select {
	case <-done:
		return
	case <-time.After(simulatedLoadTime):
	}

	select {
	case <-done:
		//fmt.Println("cancelled ... ")
		fmt.Printf("cancelled on %v\n", id)
	case result <- id:
	}

	took := time.Since(started)
	// Display how long handlers would have taken
	if took < simulatedLoadTime {
		took = simulatedLoadTime
	}
	fmt.Printf("%v took %v\n", id, took)
}

func main() {
	done := make(chan interface{})
	result := make(chan int)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go doWork(done, i, &wg, result)
	}

	firstReturned := <-result
	close(done)
	wg.Wait()

	fmt.Printf("Received an answer from #%v\n", firstReturned)
}
