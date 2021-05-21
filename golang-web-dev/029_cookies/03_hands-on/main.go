package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("counter")
	if err != nil {
		c = &http.Cookie{Name: "counter", Value: "0"}
	}

	tmp, _ := strconv.ParseInt(c.Value, 10, 32)
	c.Value = fmt.Sprint(tmp + 1)
	http.SetCookie(w, c)
	fmt.Fprintf(w, " You are the %v visitor. ", c.Value)

}
