package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

// 1. Create a data structure to pass to a template which
// * contains information about restaurant's menu including Breakfast, Lunch, and Dinner items

type meal struct {
	Name, Description string
}

type menu struct {
	Name  string
	Price int
	Meals []meal
}

type day struct {
	Day       time.Weekday
	Breakfast menu
	Lunch     menu
	Dinner    menu
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	days := []day{
		{
			Day: time.Monday,
			Breakfast: menu{
				Name:  "Breakfast",
				Price: 12,
				Meals: []meal{
					{
						Name:        "Soup",
						Description: "Yummy",
					},
					{
						Name:        "Meat",
						Description: "Tasty",
					},
				},
			},
			Lunch: menu{
				Name:  "Lunch",
				Price: 12,
				Meals: []meal{
					{
						Name:        "Soup",
						Description: "Yummy",
					},
					{
						Name:        "Meat",
						Description: "Tasty",
					},
				},
			},
			Dinner: menu{
				Name:  "Dinner",
				Price: 12,
				Meals: []meal{
					{
						Name:        "Soup",
						Description: "Yummy",
					},
					{
						Name:        "Meat",
						Description: "Tasty",
					},
				},
			},
		},
		{
			Day: time.Tuesday,
			Breakfast: menu{
				Name:  "Breakfast",
				Price: 12,
				Meals: []meal{
					{
						Name:        "Soup",
						Description: "Yummy",
					},
					{
						Name:        "Meat",
						Description: "Tasty",
					},
				},
			},
			Lunch: menu{
				Name:  "Lunch",
				Price: 12,
				Meals: []meal{
					{
						Name:        "Soup",
						Description: "Yummy",
					},
					{
						Name:        "Meat",
						Description: "Tasty",
					},
				},
			},
			Dinner: menu{
				Name:  "Dinner",
				Price: 12,
				Meals: []meal{
					{
						Name:        "Soup",
						Description: "Yummy",
					},
					{
						Name:        "Meat",
						Description: "Tasty",
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, days)
	if err != nil {
		log.Fatalln(err)
	}

}
