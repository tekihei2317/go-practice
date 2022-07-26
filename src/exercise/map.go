package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wc := make(map[string]int)

	for _, word := range words {
		wc[word]++
	}

	return wc
}

func main() {
	fmt.Println(WordCount("I am learning Go!"))
}
