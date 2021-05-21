package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.html"))
}

func main() {
	fmt.Println("The app started...")
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/index", index)
	http.HandleFunc("/aboutUs", aboutUs)
	http.HandleFunc("/contactMe", contactMe)
	http.Handle("/assets/", http.FileServer(http.Dir("./assets")))

	log.Fatalln(http.ListenAndServe(":8080", nil))
	fmt.Println("Listening on port 8080...")

}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ping action was executed...")

	fmt.Fprintln(w, "Ok.")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index action was executed...")

	err := tpl.ExecuteTemplate(w, "index.html", nil) //instance())
	handleError(err)
}

type compayInfo struct {
	Name,
	Address,
	Email,
	Phone string
}

func aboutUs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AboutUs action was executed...")

	ci := compayInfo{
		Name:    "My company",
		Address: "Cra 72 B  # 84 - 35",
		Email:   "thisIsATest@gmail.com",
		Phone:   "3002429999",
	}
	err := tpl.ExecuteTemplate(w, "aboutUs.html", ci)
	handleError(err)
}

type clientInfo struct {
	Name, Email, Phone, Reason string
}

func contactMe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ContactMe action was executed...")

	var data clientInfo

	if r.Method == http.MethodPost {
		data = clientInfo{
			Name:   r.FormValue("name"),
			Email:  r.FormValue("email"),
			Phone:  r.FormValue("phone"),
			Reason: r.FormValue("reason"),
		}
	}
	fmt.Println(data)
	err := tpl.ExecuteTemplate(w, "contactMe.html", data)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func instance() string {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	handleError(err)
	bs := make([]byte, resp.ContentLength)
	return string(bs)
}
