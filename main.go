package main

import (
	"ascii-art-web/ascii"
	"fmt"
	"html/template"
	"net/http"
)

type TemplateData struct {
	ASCIIART string
	Error    string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 PAGE NOT FOUND", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		// Serve the initial form
		err := tpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}

	case "POST":
		// Parse form data
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form data", http.StatusInternalServerError)
			return
		}

		// Get form values
		text := r.FormValue("text")
		template := r.FormValue("template")

		// Validate input
		if text == "" || template == "" {
			http.Error(w, "Text and Template fields are required", http.StatusBadRequest)
			return
		}

		// Generate ASCII Art
		asciiArt, statusCode := ascii.GenerateASCIIArt(text, template)
		if statusCode != http.StatusOK {
			http.Error(w, asciiArt, statusCode) // Send error response with proper status code
			return
		}

		// Pass ASCII Art to the template
		data := TemplateData{ASCIIART: asciiArt}
		err := tpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
