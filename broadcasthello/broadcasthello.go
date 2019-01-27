package main

import (
	"fmt"
	"sync"
)

func doer(c *sync.Cond, wg *sync.WaitGroup, id int) {
	var goroutineRunning sync.WaitGroup
	goroutineRunning.Add(1)

	go func() {
		goroutineRunning.Done()

		fmt.Println("Hello! I am", id)
		c.L.Lock()
		defer c.L.Unlock()
		c.Wait()
		fmt.Println(id, "is done :)")
		wg.Done()
	}()
	goroutineRunning.Wait()
}

/*
Quiero mostrar como un set de goroutines son notificadas para iniciar su ejecución.
*/
func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	c := sync.NewCond(&sync.Mutex{})

	doer(c, &wg, 1)
	doer(c, &wg, 2)
	doer(c, &wg, 3)

	c.Broadcast()
	wg.Wait()

}
