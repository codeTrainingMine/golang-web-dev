package main

import (
	"log"
	"os"
	"text/template"
)

func main()  {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// do not use the above code in production
// we will learn about efficiency improvements soon!