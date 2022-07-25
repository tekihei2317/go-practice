package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(36, 13))

	a, b := swap("world!", "Hello")
	fmt.Println(a, b)
}
