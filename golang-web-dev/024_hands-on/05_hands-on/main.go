package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("starting-files/templates/index.gohtml"))

}

func main() {
	http.Handle("/pics/", http.FileServer(http.Dir("./starting-files/public/pics")))
	http.HandleFunc("/", index)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	handleError(tpl.Execute(w, nil))
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
