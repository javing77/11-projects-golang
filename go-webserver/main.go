package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Form data: Recivide: \n ")

	name := r.FormValue("name")
	address := r.FormValue("address")
	phone := r.FormValue("phone")

	fmt.Fprintf(w, "Name: %s \nAddress: %s \nPhone: %s", name, address, phone)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {

	fileServer := http.FileServer(http.Dir("static/"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
