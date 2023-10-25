package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// The handler function that will be called for all requests.
	// It will load the template and execute it.
	handleIndex := func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("views/index.html"))
		template.Execute(w, nil)
	}

	handleHome := func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("views/modules/home/home.html"))
		template.Execute(w, nil)
	}

	handleLogs := func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("views/modules/logs/logs.html"))
		template.Execute(w, nil)
	}

	handleAbout := func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("views/modules/about/about.html"))
		template.Execute(w, nil)
	}

	// Associate the handler function with the root URL.
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/logs", handleLogs)
	http.HandleFunc("/about", handleAbout)
	http.HandleFunc("/home", handleHome)

	// Start the server.
	log.Println("Running on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
