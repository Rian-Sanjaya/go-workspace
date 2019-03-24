package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// ENCODING (from go type to json)

	type message struct {
		Name string `json: name`
		Body string `json: body`
		Time int64  `json: time`
	}

	m := message{"Alice", "Hello", 1294706395881547000}

	// use json.Marshal to encode m into json
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal("Error when encoding to json: ", err)
	}
	fmt.Println(string(b))

	// DECODING (from json to go type)

	// first create a place where the decoded data will be stored
	var ms message

	// then call json.Unmarshal, passing it a []byte of json data dan a pointer to m
	err1 := json.Unmarshal(b, &ms)
	if err1 != nil {
		log.Fatal("Error when decoding json: ", err)
	}
	// if b contains valid json that fits in m, the data from b will have been stored in m
	fmt.Println(m)

	// if the structure of json data doesn't exactly match the go type
	// Unmarshal will decode only the fields that it can find in the destination type.
	// In this case, only the Name field of m will be populated,
	// and the Food field will be ignored.
	// This behavior is particularly useful when you wish to pick only a few specific
	// fields out of a large JSON blob.
	// It also means that any unexported fields in the destination struct
	// will be unaffected by Unmarshal.
	var m1 message
	b1 := []byte(`{"Name": "Bob", "Food": "Pickle"}`)
	err1 = json.Unmarshal(b1, &m1)
	fmt.Println(m1.Name, ",", m1.Body, ",", m1.Time)

	// what if we don't know the structure of the json data beforehand?
	// The interface{} (empty interface) type describes an interface with zero methods.
	// Every Go type implements at least zero methods and therefore satisfies the empty interface.
	// The empty interface serves as a general container type

}
