package main

import (
	"fmt"
	"math"
)

// By using ... before the type name of the last parameter you can indicate that it takes zero or more of those parameters
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	
	return total
}

// Functions are values too. They can be passed around just like other values.
// Function values may be used as function arguments and return values.
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	fmt.Println("\nVARIADIC FUNCTION (... PARAMETER)")
	fmt.Println(add(1,2,3,4,5))
	
	xs := []int{1,2,3,4,5}
	
	// pass a slice of ints by following the slice with ...
	fmt.Println(add(xs...))

	fmt.Println("\nFUNCTION AS VALUE")
	// Functions are values too. They can be passed around just like other values.
	// Function values may be used as function arguments and return values.
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
}