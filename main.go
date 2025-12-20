package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler function for /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	//!handle the path is not what we expected it to be (!= "/hello")
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	//!handle someone tries a GET method (dispalys hello only)
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

// handler function for /form
func formHandler(w http.ResponseWriter, r *http.Request) {
	//!handle parsing error of the form
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")

	//*handle form values comming from the user
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
}

func main() {
	//create file server
	fileServer := http.FileServer(http.Dir("./static"))
	//handle the root route
	http.Handle("/", fileServer)
	//handle /form route
	http.HandleFunc("/form", formHandler)
	//handle /hello route
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port :8080")

	//!handle the port not listening and log the error
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
