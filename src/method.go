package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x float64
	y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Scale(v Vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v, v.Abs())

	v.Scale(10)
	fmt.Println(v, v.Abs())

	x := Vertex{5, 12}
	Scale(x, 10)
	fmt.Println(x)
}
