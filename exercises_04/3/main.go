// 1. Take the previous program and change it so that:
// * func main uses http.Handle instead of http.HandleFunc
//
// Contstraint: Do not change anything outside of func main
//
// Hints:
//
// [http.HandlerFunc](https://godoc.org/net/http#HandlerFunc)
// ``` Go
// type HandlerFunc func(ResponseWriter, *Request)
// ```
//
// [http.HandleFunc](https://godoc.org/net/http#HandleFunc)
// ``` Go
// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
// ```
//
// [source code for HandleFunc](https://golang.org/src/net/http/server.go#L2069)
// ``` Go
//   func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
//   		mux.Handle(pattern, HandlerFunc(handler))
//   }
// ```
package main

import (
	"html/template"
	"log"
	"net/http"
)

func e(res http.ResponseWriter, req *http.Request) {
	execute(res, "empty")
}

func d(res http.ResponseWriter, req *http.Request) {
	execute(res, "dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	execute(res, "cat cat")
}

func m(res http.ResponseWriter, req *http.Request) {
	execute(res, "chris")
}

func execute(res http.ResponseWriter, value string) {
	err := tpl.Execute(res, value)
	if err != nil {
		log.Fatalln(err)
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(e))
	http.Handle("/dog", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))
	http.Handle("/me/", http.HandlerFunc(m))

	http.ListenAndServe(":8080", nil)
}
