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


func renderError(w http.ResponseWriter, status int, errorTemplate string) {
	w.WriteHeader(status)
	err := tpl.ExecuteTemplate(w, errorTemplate, nil)
	if err != nil {
		http.Error(w, http.StatusText(status), status)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound, "404.html")
		return
	}

	switch r.Method {
	case "GET":
		// Serve the initial form
		err := tpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			renderError(w, http.StatusInternalServerError, "500.html")
		}

	case "POST":
		// Parse form data
		if err := r.ParseForm(); err != nil {
			renderError(w, http.StatusBadRequest, "400.html") // Handle form parsing errors
			return
		}

		// Get form values
		text := r.FormValue("text")
		template := r.FormValue("template")

		// Validate input
		if text == "" || template == "" {
			// Call renderError for 400.html on validation failure
			renderError(w, http.StatusBadRequest, "400.html")
			return
		}

		// Generate ASCII Art
		asciiArt, statusCode := ascii.GenerateASCIIArt(text, template)
		if statusCode != http.StatusOK {
			// Handle ASCII art generation errors
			switch statusCode {
			case http.StatusBadRequest:
				renderError(w, http.StatusBadRequest, "400.html")
			case http.StatusInternalServerError:
				renderError(w, http.StatusInternalServerError, "500.html")
			default:
				http.Error(w, "Unexpected error", statusCode)
			}
			return
		}

		// Pass ASCII Art to the template
		data := TemplateData{ASCIIART: asciiArt}
		err := tpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			renderError(w, http.StatusInternalServerError, "500.html")
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



