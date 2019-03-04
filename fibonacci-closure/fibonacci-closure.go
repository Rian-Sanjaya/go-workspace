package main

import (
	"fmt"
)

func fibonacci() func() int {
	a := 1
	b := 0

	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// a  b
// 0  1
// 1  1
// 1  2
// 2  3
// 3  5
// 5  8
