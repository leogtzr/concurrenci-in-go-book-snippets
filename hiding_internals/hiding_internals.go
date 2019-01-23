package main

import "sync"

// Counter ...
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment ...
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func main() {

}
