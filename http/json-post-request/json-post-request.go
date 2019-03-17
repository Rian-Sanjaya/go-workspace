package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type testStruct struct {
	Test string
}

func test(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t testStruct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.Test)
}

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
