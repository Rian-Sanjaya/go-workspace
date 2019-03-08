package main

import "fmt"

func main() {
	// declare a slice
	var x []float64

	// create a slice, with size of 5, and capacity of 10
	x = make([]float64, 5, 10)
	fmt.Println(x)

	// another way to create slice is to use: [low : high]
	// low is the index where to start and high is the index where to end (not including the index itself)
	arr := [5]float64{1, 2, 3, 4, 5}
	sl1 := arr[0:5]
	sl2 := arr[0:]
	sl3 := arr[:5]
	fmt.Println(sl1, sl2, sl3)

	// append creates a new slice by taking an existing slice (the first argument) and appending all the following arguments to it
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	// the contents of slice3 are copied into slice4,
	// but since slice4 has room for only two elements only the first two elements of slice3 are copied
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
}
