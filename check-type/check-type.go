package main

import (
	"fmt"
	"reflect"
)

func main() {
	sl := []int{3, 5}
	fmt.Printf("%T\n", sl)
	fmt.Println(reflect.TypeOf(sl))
}
