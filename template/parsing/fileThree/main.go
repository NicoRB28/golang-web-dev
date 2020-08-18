package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	//Glob me parsea todo lo que responda al patron
	tpl, err := template.ParseGlob("../fileTwo/*.gmao")

	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
