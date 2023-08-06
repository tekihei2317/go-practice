package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func structExample() {
	v := Vertex{1, 2}
	fmt.Println(v)

	v.X = 4
	fmt.Println(v)
}

func arrayExample() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println(primes)
}

func sliceExample() {
	names := [4]string{
		"Taro",
		"Hanako",
		"Yamada",
		"Sato",
	}
	s := names[1:3]

	fmt.Println(names, s)

	names[1] = "Update"
	fmt.Println(names, s)

	s[0] = "Update2"
	fmt.Println(names, s)
}

func sliceIteration() {
	primes := []int{2, 3, 5, 7, 11, 13}

	for i, v := range primes {
		fmt.Printf("i = %d, v = %d\n", i, v)
	}
}

func mapExample() {
	var m = map[int]string{
		1: "hello",
		2: "world",
	}

	m[3] = "test"
	fmt.Println(m)
}

func mapMutation() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m)

	delete(m, "Answer")
	fmt.Println("The value:", m)

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func main() {
	structExample()
	arrayExample()
	sliceExample()
	sliceIteration()
	mapExample()
	mapMutation()
}
