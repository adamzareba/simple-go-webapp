package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
)

type AboutAppPage struct {
	Title       string
	Description string
}

func main() {
	http.HandleFunc("/", home_handler)
	http.HandleFunc("/about", about_app_handler)
	log.Println(http.ListenAndServe(":8080", nil))
}

func home_handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello my user!")
}

func about_app_handler(writer http.ResponseWriter, request *http.Request) {
	page := AboutAppPage{Title: "Simple Go WebApp", Description: "Test application for demo purpose"}
	t, _ := template.ParseFiles("about_app.html")
	t.Execute(writer, page)
}
