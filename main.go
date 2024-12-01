package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// TemplateData holds data to be passed to the HTML template
type TemplateData struct {
	ASCIIART string
	Error    string
}

var tpl *template.Template

func init() {
	// Parse all templates in the "templates" folder
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

		// Generate ASCII Art (mock function, replace with actual implementation)
		asciiArt, err := generateASCIIArt(text, template)
		if err != nil {
			data := TemplateData{Error: err.Error()}
			tpl.ExecuteTemplate(w, "index.html", data)
			return
		}

		// Pass ASCII Art to the template
		data := TemplateData{ASCIIART: asciiArt}
		err = tpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func generateASCIIArt(text, template string) (string, error) {

	if template != "standard" && template != "shadow" && template != "thinkertoy" {
		return "", fmt.Errorf("invalid template selected")
	}

	return fmt.Sprintf("Generated ASCII art for '%s' using '%s' template", text, template), nil
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}



