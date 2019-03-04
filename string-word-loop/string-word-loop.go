package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "This, that, and the other"

	words := strings.Fields(s)
	// fmt.Println(words)
	for i, word := range words {
		fmt.Println(i, " => ", word)
	}

	// replace punctuation
	replacer := strings.NewReplacer(",", "", ".", "", ";", "")
	st := replacer.Replace(s)
	ws := strings.Fields(st)
	for i, w := range ws {
		fmt.Println(i, " => ", w)
	}

	// replace punctuation using map
	removePunctuation := func(r rune) rune {
		if strings.ContainsRune(".,:;", r) {
			return -1
		}

		return r
	}

	str := strings.Map(removePunctuation, s)
	wrds := strings.Fields(str)
	for i, w := range wrds {
		fmt.Println(i, " => ", w)
	}

	s = "This, that, other"
	words = strings.Split(s, ", ")
	for i, word := range words {
		fmt.Println(i, " => ", word)
	}
}
