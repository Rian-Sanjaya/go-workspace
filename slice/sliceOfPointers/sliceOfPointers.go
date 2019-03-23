package main

import (
	"fmt"
	"strconv"
)

func main() {
	// declare empty slice of string pointer
	sl := []*string{}

	for i := 0; i < 10; i++ {
		// a new object s will be create in a new memory address every loop
		var s string
		// convert int to string
		s = fmt.Sprintf("#%s", strconv.Itoa(i))
		// append to slice a pointer / memory address of s
		sl = append(sl, &s)
	}

	fmt.Printf("%v\n", sl)
	for _, v := range sl {
		// get the value from the memory address v pointing to
		fmt.Printf("%v\n", *v)
	}

}
