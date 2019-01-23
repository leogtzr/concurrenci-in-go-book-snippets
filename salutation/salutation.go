package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	saludo := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		saludo = "holis"
	}()

	wg.Wait()
	fmt.Println(saludo)
}
