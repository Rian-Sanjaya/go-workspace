package main

import "fmt"

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
}

func main() {
	fmt.Println("\nSLICE LITERAL")
	// Array literal
	//[3]bool{true, true, false}
	// Slice literal
	// this creates the same array as above, then builds a slice that references it
	//[]bool{true, true, false}

	// 1. CREATE A SLICE
	s0 := []string{"one", "two", "three"}
	fmt.Println(s0)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
	}
	fmt.Println(s)

	fmt.Println("SLICE LENGTH AND CAPACITY")
	// 2. Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	// The make function allocates a zeroed array and returns a slice that refers to that array
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
	x := make([]int, 0, 5) // create slice 0 length and 5 capacity
	printSlice(x)

	x1 := []int{2, 3, 5, 7, 11, 13}
	printSlice(x1)

	// slice the slice to give it zero length
	x1 = x1[:0]
	printSlice(x1)

	// extend its length
	x1 = x1[:4]
	printSlice(x1)

	// drop its first two values
	x1 = x1[2:]
	printSlice(x1)

	fmt.Println("\nSLICE DEFAULT")
	// 3. another way to create slice is to use: [low : high]
	// low is the index where to start and high is the index where to end (not including the index itself)
	arr := [5]float64{1, 2, 3, 4, 5}
	sl1 := arr[0:5]
	sl2 := arr[0:]
	sl3 := arr[:5]
	fmt.Println(sl1, sl2, sl3)

	fmt.Println("\nSLICE ARE LIKE REFERENCE TO ARRAY")
	// Slices are like references to arrays
	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	fmt.Println("\nNIL SLICE")
	// NIL SLICE
	// The zero value of a slice is nil
	// A nil slice has a length and capacity of 0 and has no underlying array
	var ns []int
	fmt.Println(s, len(ns), cap(ns))
	if ns == nil {
		fmt.Println("nil slice")
	}

	fmt.Println("\nAPPEND SLICE")
	// append creates a new slice by taking an existing slice (the first argument) and appending all the following arguments to it
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	fmt.Println("\nCOPY SLICE")
	// the contents of slice3 are copied into slice4,
	// but since slice4 has room for only two elements only the first two elements of slice3 are copied
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

	fmt.Println("\nDYNAMIC SLICE")
	var slc []int
	printSlice(slc)

	// append works on nil slices
	slc = append(slc, 0)
	printSlice(slc)

	// the slice grow as needed
	slc = append(slc, 1)
	printSlice(slc)

	// we can add more than one element at a time
	slc = append(slc, 2, 3, 4)
	printSlice(slc)
}
