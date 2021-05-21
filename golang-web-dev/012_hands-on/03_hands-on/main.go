package main

import (
	"fmt"
	"html/template"
	"os"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("tpl.html"))
}

func main() {
	err := t.Execute(os.Stdout, californiaHotels)
	if err != nil {
		fmt.Println(err)
	}
}

type hotel struct {
	Name, Address, City, Zip, Region string
}

var californiaHotels = []hotel{
	hotel{
		Name:    "test2",
		Address: "Cra 72 B",
		City:    "City 1",
		Zip:     "05001",
		Region:  "Southern",
	},
	hotel{
		Name:    "test3",
		Address: "Cra 73 B",
		City:    "City 3",
		Zip:     "05003",
		Region:  "Southern",
	},
	hotel{
		Name:    "test4",
		Address: "Cra 74 B",
		City:    "City 4",
		Zip:     "05004",
		Region:  "Central",
	},
	hotel{
		Name:    "test5",
		Address: "Cra 75 B",
		City:    "City 5",
		Zip:     "05005",
		Region:  "Central",
	},
}
