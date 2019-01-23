package main

import "fmt"

func main() {
	go hello()

	go func() {
		fmt.Println("world")
	}()

}

func hello() {
	fmt.Println("hello")
}
