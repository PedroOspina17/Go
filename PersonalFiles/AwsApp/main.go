package main

import (
	"fmt"
	"io/ioutil"
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
	http.HandleFunc("/", index)
	http.HandleFunc("/aboutUs", aboutUs)
	http.HandleFunc("/contactMe", contactMe)
	http.Handle("/assets/", http.FileServer(http.Dir("./assets")))

	log.Fatalln(http.ListenAndServe(":80", nil))
	fmt.Println("Listening on port 8080...")

}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ping action was executed...")

	fmt.Fprintln(w, "Ok.")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index action was executed...")

	err := tpl.ExecuteTemplate(w, "index.html", instance())
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
	fmt.Println("Looking for the instance Id...")
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	handleError(err)
	defer resp.Body.Close()

	data, er := ioutil.ReadAll(resp.Body)
	handleError(er)

	// Another option can be:
	// data2 := make([]byte, resp.ContentLength)
	// resp.Body.Read(data2)

	fmt.Println("Instance Id: ", string(data))
	return string(data)
}
