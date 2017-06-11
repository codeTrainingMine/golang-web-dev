package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name string
	Motto string
}

type car struct {
	Manufacturer string
	Model string
	Doors int
}

type items struct {
	Wisdom []sage
	Transport []car
}

func init()  {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main()  {

	b := sage {
		Name: "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name: "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name: "Martin Luther Kind",
		Motto: "Hatred never ceases with hatred but with love alond is healed.",
	}

	f := car {
		Manufacturer:"Ford",
		Model:"F150",
		Doors: 2,
	}

	bm := car {
		Manufacturer:"BMW",
		Model:"M235",
		Doors: 2,
	}

	sages := []sage{b, g, m}
	cars := []car{f, bm}

	data := items{
		Wisdom: sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}