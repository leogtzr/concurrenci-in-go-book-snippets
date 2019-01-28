package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	instance := myPool.Get() // New is invoked here....
	myPool.Put(instance)     // We put the instance here available for later re-use.
	myPool.Get()

}
