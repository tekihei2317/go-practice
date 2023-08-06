package main

import (
	"fmt"
)

func main() {
	sum := 0

	for i := 1; i <= 100; i++ {
		sum += i
	}

	sum2 := 0
	for sum2 < 100 {
		sum2++
	}

	fmt.Println(sum, sum2)
}
