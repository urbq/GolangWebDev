package main

import (
	"log"
	"os"
	"text/template"
)

// 1. Create a data structure to pass to a template which
// * contains information about California hotels including Name, Address, City, Zip, Region
// * region can be: Southern, Central, Northern
// * can hold an unlimited number of hotels

type hotel struct {
	Name, Address, City, Zip string
}

type region struct {
	Name   string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// hotels := []hotel{
	// 	hotel{"ABC", "ADDRESS", "CITY", "ZIP", "REGION"},
	// 	hotel{"abc", "address", "city", "zip", "region"},
	// }
	regions := []region{
		region{"SOUTH",
			[]hotel{
				{"ABC", "ADDRESS", "CITY", "ZIP"},
				{"abc", "address", "city", "zip"},
			},
		},
		region{"NORTH",
			[]hotel{
				{"ABC", "ADDRESS", "CITY", "ZIP"},
				{"abc", "address", "city", "zip"},
			},
		}}
	err := tpl.Execute(os.Stdout, regions)
	if err != nil {
		log.Fatalln(err)
	}

}
