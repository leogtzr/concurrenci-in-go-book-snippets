package main

import (
	"fmt"
	"sync"
	"time"
)

// Record ...
type Record struct {
	sync.Mutex
	data string
	cond *sync.Cond
}

// NewRecord ....
func NewRecord() *Record {
	r := Record{}
	r.cond = sync.NewCond(&r)
	return &r
}

func main() {
	var wg sync.WaitGroup
	rec := NewRecord()
	wg.Add(1)

	fmt.Println("Starting up ... ")

	/*
		It may look like that the goroutine waiting(rec.cond.Wait()) is holding the lock whole time(rec.Lock()), but its not,
		Internally cond.Wait() unlocks it and it locks it again only when it wakes up by other go routine.
	*/

	go func(rec *Record) {
		defer wg.Done()
		rec.Lock()
		rec.cond.Wait()
		rec.Unlock()
		fmt.Println("Unlocked ... ")
		fmt.Println("Data: ", rec.data)
		return
	}(rec)

	time.Sleep(2 * time.Second)
	rec.Lock()
	rec.data = "leonardo"
	rec.Unlock()
	rec.cond.Signal()

	fmt.Println("Bye ... ")

	wg.Wait()
}
