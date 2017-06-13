package main

import (
	"os"
	"log"
	"text/template"
)

type meal struct {
	Name string
	Price float64
}

type menu struct {
	Breakfast []meal
	Lunch []meal
	Dinner []meal
}

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main()  {
	menu1 := menu{
		Breakfast: []meal{
			{Name: "Scrambled Eggs", Price: 11.50},
			{Name: "Eggs & Bacon", Price: 15.25},
			{Name: "Surfer's Burger", Price: 17.50},
		},
		Lunch: []meal{
			{Name: "Toasted Foccaccia", Price: 19.50},
			{Name: "Rigatoni Napoli", Price: 21.25},
			{Name: "Steak Sandwich", Price: 22.50},
		},
		Dinner: []meal{
			{Name: "Steak & Potatoes", Price: 30.50},
			{Name: "Salmon & Chips", Price: 27.25},
			{Name: "Mushrom Rissoto", Price: 26.50},
		},
	}

	err := tpl.Execute(os.Stdout, menu1)
	if err != nil {
		log.Fatalln(err)
	}
}

//1. Create a data structure to pass to a template which
//* contains information about restaurant's menu including Breakfast, Lunch, and Dinner items