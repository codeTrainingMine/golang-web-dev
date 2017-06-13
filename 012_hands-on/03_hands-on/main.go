package main

import (
	"os"
	"log"
	"text/template"
)


type hotel struct {
	Name, Address, City, Zip string
}

type hotels struct {
	Southern []hotel
	Central []hotel
	Northern []hotel
}

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main()  {
	h := hotels{
		Southern: []hotel{
			{Name: "Hilton", Address: "12 Broadway", City:"Houston", Zip:"31312"},
			{Name: "Parkside", Address: "33 High Street", City:"Florida", Zip:"54543"},
			{Name: "Holday Inn", Address: "3 Main Road", City:"Dallas", Zip:"45564"},
		},
		Central: []hotel{
			{Name: "Hilton", Address: "12 Broadway", City:"Los Angeles", Zip:"31312"},
			{Name: "Parkside", Address: "33 High Street", City:"San Francisco", Zip:"54543"},
			{Name: "Holday Inn", Address: "3 Main Road", City:"Boston", Zip:"45564"},
		},
		Northern: []hotel{
			{Name: "Hilton", Address: "12 Broadway", City:"New York", Zip:"31312"},
			{Name: "Parkside", Address: "33 High Street", City:"Washington", Zip:"54543"},
			{Name: "Holday Inn", Address: "3 Main Road", City:"Redmond", Zip:"45564"},
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}

//1. Create a data structure to pass to a template which
//* contains information about California hotels including Name, Address, City, Zip, Region
//* region can be: Southern, Central, Northern
//* can hold an unlimited number of hotels