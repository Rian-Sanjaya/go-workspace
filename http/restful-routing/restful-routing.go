package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", HomeHandler)

	r.GET("/posts", PostsIndexHandler)
	r.POST("/posts", PostCreateHandler)

	r.GET("/posts/:id", PostShowHandler)
	r.PUT("/posts/:id", PostUpdateHandler)
	r.GET("/posts/:id/edit", PostEditHandler)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}

func HomeHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintln(res, "Home")
}

func PostsIndexHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintln(res, "posts index")
}

func PostCreateHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintln(res, "posts create")
}

func PostShowHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Fprintln(res, "showing post", id)
}

func PostUpdateHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintln(res, "post update")
}

func PostDeleteHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintln(res, "post delete")
}

func PostEditHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintln(res, "post edit")
}
