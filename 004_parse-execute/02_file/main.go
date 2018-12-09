package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// getting pointer to tpl
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating fie", err)
	}

	// now with pointer to templat you can Execute
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
