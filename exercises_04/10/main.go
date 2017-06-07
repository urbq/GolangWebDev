// # Building upon the code from the previous problem:
//
// Add code to respond to the following METHODS & ROUTES:
// 	GET /
// 	GET /apply
// 	POST /apply

package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	mux := httprouter.New()
	mux.GET("/", defaultGet)
	mux.GET("/apply", applyGet)
	mux.POST("/apply", applyPost)

	http.ListenAndServe(":8080", mux)
}

func defaultGet(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	io.WriteString(w, `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
		</head>
		<body>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
			<form action="/apply" method="POST">
			<input type="hidden" value="In my good death">
			<input type="submit" value="submit">
			</form>
		</body>
		</html>
	`)
}

func applyGet(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	io.WriteString(w, "applyGet")
}

func applyPost(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	io.WriteString(w, "applyPost")
}
