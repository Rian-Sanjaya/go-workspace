package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)

		// f prints out the numbers from 0 to 10, waiting between 0 and 250 ms after each one
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}

	var input string
	fmt.Scanln(&input)
	// call to the Scanln function has been included;
	// without it the program would exit before being given the opportunity to print all the numbers
}
