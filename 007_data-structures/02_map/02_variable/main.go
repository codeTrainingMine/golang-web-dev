package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main()  {

	sages := map[string]string{
		"India": "Gandhi",
		"America": "MLK",
		"Meditate": "Budhha",
		"Love": "Jesus",
		"Prophet": "Muhammad",
	}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}