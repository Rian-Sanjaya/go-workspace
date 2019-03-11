package main

import (
	"fmt"
)

func main() {
	// 1. create map
	m := make(map[string]int)
	m["one"] = 1
	fmt.Println(m["two"])

	// 2. create map
	m1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(m1)

	// check if key exists
	v, ok := m1["three"]
	fmt.Println("value:", v, "key:", ok)

	if _, ok := m1["three"]; ok {
		fmt.Println("key exists")
	} else {
		fmt.Println("key not exists")
	}

	// delete a key
	delete(m1, "two")
	fmt.Println(m1)
	m1["three"] = 3
	fmt.Println(m1)
}
