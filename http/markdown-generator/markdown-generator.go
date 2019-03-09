package main

import (
	"net/http"

	"github.com/russross/blackfriday"
)

func main() {
	// handle the '/markdown' route
	// http.HandleFunc can be think of handling routes via a function instead of an object
	http.HandleFunc("/markdown", GenerateMarkDown)

	// serve the public (index.html)
	// http.Handle on the '/' pattern will catch-all-route, so we define that route last
	// http.FileServer returns an http.Handler
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":8080", nil)
}

// renders Html from a form field containing markdown-formatted text
func GenerateMarkDown(res http.ResponseWriter, req *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(req.FormValue("body")))
	res.Write(markdown)
}
