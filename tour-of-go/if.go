package main

import "fmt"

func isEven(n int) bool {
	if rem := n % 2; rem == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isEven(3), isEven(10))
}
