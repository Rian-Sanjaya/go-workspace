package main

import "net/http"

func main() {
	// http.ListenAndServe function is used to start the server
	// http.FileServer will server the current working directory with http.Dir(".")
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
