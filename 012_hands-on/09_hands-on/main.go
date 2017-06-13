package main

import (
	//"fmt"
	"os"
	"log"
	"bufio"
	"encoding/csv"
	"io"
	"text/template"
)

type row struct {
	Date, Open, High, Low, Close, Volume, AdjClose string
}

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main()  {
	file1, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file1.Close()

	r := csv.NewReader(bufio.NewReader(file1))

	var rows []row

	var counter int
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if counter > 0 {
			rows = append(rows, row{
				Date: record[0],
				Open: record[1],
				High: record[2],
				Low: record[3],
				Close: record[4],
				Volume: record[5],
				AdjClose: record[6],
			})
		}

		counter++
	}

	//for _, value := range rows {
	//	fmt.Println(value)
	//}
	err = tpl.Execute(os.Stdout, rows)
	if err != nil {
		log.Fatalln(err)
	}
}

//Date,Open,High,Low,Close,Volume,Adj Close
//1. Parse this CSV file, putting two fields from the contents of the CSV file into a data structure.
//2. Parse an html template, then pass the data from step 1 into the CSV template; have the html template display the CSV data in a web page.