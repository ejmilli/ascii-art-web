package main

import (
	"ascii-art-web/ascii"
	"fmt"
	"html/template"
	"net/http"
)

type TemplateData struct {
	ASCIIART string
}

var tpl *template.Template

// renderError handles error page rendering
func renderError(w http.ResponseWriter, status int, errorTemplate string) {
	// Set the response status code and render the appropriate error template
	w.WriteHeader(status)
	err := tpl.ExecuteTemplate(w, errorTemplate, nil)
	if err != nil {
		// If rendering the error template fails, send a generic error
		http.Error(w, http.StatusText(status), status)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Check for invalid URL path
	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound, "404.html")
		return
	}

	switch r.Method {
	case "GET":
		// Serve the initial form on GET request
		err := tpl.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			renderError(w, http.StatusInternalServerError, "500.html")
		}

	case "POST":
		// Parse form data on POST request
		if err := r.ParseForm(); err != nil {
			renderError(w, http.StatusBadRequest, "400.html")
			return
		}

		// Get form values
		text := r.FormValue("text")
		template := r.FormValue("template")

		// Validate input
		if text == "" || template == "" {
			renderError(w, http.StatusBadRequest, "400.html")
			return
		}

		// Generate ASCII Art
		asciiArt, statusCode := ascii.GenerateASCIIArt(text, template)
		if statusCode != http.StatusOK {
			renderError(w, statusCode, fmt.Sprintf("%d.html", statusCode))
			return
		}

		// Pass ASCII Art to the template
		data := TemplateData{ASCIIART: asciiArt}
		err := tpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			renderError(w, http.StatusInternalServerError, "500.html")
		}
	}
	
}

func main() {

	tpl = template.Must(template.ParseGlob("template/*.html"))

	http.HandleFunc("/", handler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
