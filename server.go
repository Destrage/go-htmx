package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// The handler function that will be called for all requests.
	// It will load the template and execute it.
	handleHome := func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("views/index.html"))
		template.Execute(w, nil)
	}

	handleClicked := func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("views/clicked.html"))
		template.Execute(w, nil)
	}

	// Associate the handler function with the root URL.
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/clicked", handleClicked)

	// Start the server.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
