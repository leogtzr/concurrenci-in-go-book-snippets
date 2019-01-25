package main

import (
	"fmt"
	"sync"
)

func main() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v\n", id)
	}

	const numOfGreeters = 5
	var wg sync.WaitGroup

	wg.Add(numOfGreeters)

	for i := 0; i < numOfGreeters; i++ {
		go hello(&wg, i)
	}

	wg.Wait()

	fmt.Println("Done ... ")

}
