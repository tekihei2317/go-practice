package main

import (
	"fmt"
)

func main() {
	i := 42
	p := &i

	*p = 21

	fmt.Println("i =", i, ",", "*p =", *p)
}
