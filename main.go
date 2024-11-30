package main

import (
	"html/template"
	"net/http"
)


var tpl *template.Template

func init() {
	// Parse the template during initialization so it's available globally
	tpl = template.Must(template.ParseFiles("templates/index.html"))}

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	// Render the template when the endpoint is hit
	err := tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
func main() {



	http.HandleFunc("/", HandleFunc)
	 http.ListenAndServe(":8080", nil)
}


