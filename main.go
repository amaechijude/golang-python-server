package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supportted", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	email := r.FormValue("email")
	age := r.FormValue("age")
	comments := r.FormValue("comments")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Email = %s\n", email)
	fmt.Fprintf(w, "Age = %s\n", age)
	fmt.Fprintf(w, "Comments = %s", comments)
}
func balHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/bal_sheet.html")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func main() {
	//Specify the directory of the static files to be served
	fileServer := http.FileServer(http.Dir("./static"))

	//Serve the index file
	http.Handle("/", fileServer)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/bal", balHandler)
	//Indicate the form handler
	http.HandleFunc("/form", formHandler)
	//Indicate the hello handler
	http.HandleFunc("/hello", helloHandler)
	//Start the server on port 8080
	fmt.Printf("Server Started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
