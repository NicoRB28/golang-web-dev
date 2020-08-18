package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 28)
	//el 28 va a suplantar el "." dentro de los {{}} en el template
	//si bien solo se puede pasar un solo dato, ese dato puede ser una agregación es decir un map, un slice
	//eso significa que en realidad puedo pasar varios datos si lo hago correctamente aprovechando las
	//estructuras de datos mas complejas
	if err != nil {
		log.Fatalln(err)
	}

}
