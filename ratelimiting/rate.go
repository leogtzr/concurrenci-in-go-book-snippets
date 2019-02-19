package main

import (
	"context"
	"log"
	"os"
	"sync"
)

func main() {
	defer log.Printf("Done ... main done ... ")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("cannot ReadFile: %v\n", err)
			}
			log.Printf("ReadFile")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("cannot ResolveAddress: %v\n", err)
			}
			log.Printf("ResolveAddress")
		}()
	}

	wg.Wait()
1. m. Nota de burla o de afrenta. Le puso el inri.

para más, o mayor, inri

1. locs. advs. Para mayor escarnio.
}
