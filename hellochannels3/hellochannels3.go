package main

import "fmt"

func main() {

	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello world"
	}()

	if salutation, ok := <-stringStream; ok {
		fmt.Printf("Ok (%v)\n", salutation)
	}

}
