package main

import (
	"fmt"
	"math"
)

// 1. declare a struct
type circle struct {
	x, y, r float64
}

func main() {
	// 2. create a struct
	// c := new(circle) // This allocates memory for all the fields, sets each of them to their zero value and returns a pointer
	// fmt.Println(c)
	// c = circle{x: 3, y: 4, r: 5}
	c := circle{0, 0, 5}
	fmt.Println(circleArea(c))
}

func circleArea(c circle) float64 {
	return math.Pi * c.r * c.r
}
