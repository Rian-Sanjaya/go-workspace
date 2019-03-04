package main

import "fmt"

func main() {
	for i, char := range "aku seorang kapiten" {
		fmt.Printf("Char %c is at position %d\n", char, i)
	}

	for i, c := range "abc" {
		fmt.Println(i, " => ", string(c))
	}
}
