package main

import (
	"log"
	"os"
	"text/template"
)

type car struct {
	Name   string
	Engine string
	TaxDeduct   bool
}

func main() {
	c := car{
		Name:   "Tesla",
		Engine: "Electric",
		TaxDeduct:   false,
	}

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, c)
	if err != nil {
		log.Fatalln(err)
	}
}