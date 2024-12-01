package main

import (
	"ascii-art-web/ascii"
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

type TemplateData struct {
	ASCIIART string
}


func init() {
	var err error

	tpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
			fmt.Println("Error loading templates:", err)
			return
	}
	fmt.Println("Templates loaded successfully.")
}



func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Render the template with an empty result
		err := tpl.ExecuteTemplate(w, "index.html", TemplateData{})
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
		return
	}

	if r.Method == "POST" {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusInternalServerError)
			return
		}

		// Get text and template from the form
		text := r.FormValue("text")
		templateName := r.FormValue("template")

		// Generate ASCII art using your ASCII package
		asciiArt, err := ascii.GenerateASCIIArt(text, templateName) // Replace with your actual function
		if err != nil {
			http.Error(w, "Error generating ASCII art: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Pass the ASCII art back to the template
		data := TemplateData{ASCIIART: asciiArt}
		err = tpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
