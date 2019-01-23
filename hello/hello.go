package main

import "fmt"

func main() {
	go hello()

	go func() {
		fmt.Println("world")
	}()

	hi := func() {
		fmt.Println("holis")
	}

	go hi()

}

func hello() {
	fmt.Println("hello")
}
