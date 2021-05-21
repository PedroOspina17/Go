package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}
func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text-html")
	fmt.Fprintln(w, "<h1>This is the index action</h1>")
}

func dog(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "dog.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err.Error())
	}
}
