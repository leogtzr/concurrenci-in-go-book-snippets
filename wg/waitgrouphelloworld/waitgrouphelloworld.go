package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping ... ")
		time.Sleep(1 * time.Second)
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("2st goroutine sleeping ... ")
		time.Sleep(2 * time.Second)
	}()

	wg.Wait()
	fmt.Println("All gouroutines complete.")
}
