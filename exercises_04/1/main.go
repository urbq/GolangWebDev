// ListenAndServe on port ":8080" using the default ServeMux.
//
// Use HandleFunc to add the following routes to the default ServeMux:
//
// "/"
// "/dog/"
// "/me/
//
// Add a func for each of the routes.
//
// Have the "/me/" route print out your name.
package main

import (
	"io"
	"net/http"
)

func e(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "empty route")
}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func m(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "chris")
}

func main() {

	http.HandleFunc("/", e)
	http.HandleFunc("/dog", d)
	http.HandleFunc("/me/", m)

	http.ListenAndServe(":8080", nil)
}
