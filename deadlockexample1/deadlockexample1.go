package main

import (
	"fmt"
	"sync"
	"time"
)

type Value struct {
	mu    sync.Mutex
	value int
}

var wg sync.WaitGroup

func main() {
	printSum := func(v1, v2 *Value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b Value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&a, &b)

	wg.Wait()
}
