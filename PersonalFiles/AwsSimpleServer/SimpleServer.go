package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	log.Fatalln(http.ListenAndServe(":80", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world !! This is my first AWS Go App.\n\n\n Developed by: Pedro Ospina.")

}
