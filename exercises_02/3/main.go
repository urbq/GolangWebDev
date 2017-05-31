package main

import (
	"encoding/csv"
	"log"
	"os"
	"text/template"
)

// 1. Parse this CSV file, putting two fields from the contents of the CSV file into a data structure.
// 2. Parse an html template, then pass the data from step 1 into the CSV template; have the html template display the CSV data in a web page.

type item struct {
	Date, Open string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.Execute(os.Stdout, parse("table.csv"))
	if err != nil {
		log.Fatalln(err)
	}

}

func parse(path string) []item {
	file, _ := os.Open(path)
	defer file.Close()
	r := csv.NewReader(file)
	lines, _ := r.ReadAll()
	// var data []item
	data := make([]item, 0, len(lines))

	for i := 1; i < len(lines); i++ {
		data = append(data, item{lines[i][0], lines[i][1]})
	}

	return data
}
