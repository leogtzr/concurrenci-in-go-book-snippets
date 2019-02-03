package main

import "fmt"

func xs(x int, nums ...int) int {
	return nums[x]
}

func main() {
	fmt.Println(xs(1, 3, 5, 7))
}
