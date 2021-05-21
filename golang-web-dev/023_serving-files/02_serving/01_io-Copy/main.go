package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	fmt.Println("dog function was executed...")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="/toby.jpg">
	`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	fmt.Println("dogpic function was executed...")
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
