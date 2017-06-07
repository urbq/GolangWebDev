// 1. Take the previous program in the previous folder and change it so that:
// * a template is parsed and served
// * you pass data into the template
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

	http.HandleFunc("/", e)
	http.HandleFunc("/dog", d)
	http.HandleFunc("/me/", m)

	http.ListenAndServe(":8080", nil)
}
