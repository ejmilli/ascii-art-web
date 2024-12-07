package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response struct to format messages
type Response struct {
	Message string `json:"message"`
}

// Enable CORS middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Function to generate ASCII Art (based on the template selected)
func generateASCII(template string, text string) string {
	// This is just a placeholder. Replace it with your ASCII generation logic.
	// Depending on the template selected, you load the corresponding map/template
	asciiMap := map[string]string{
		"standard":   "ASCII Art Standard: " + text,
		"shadow":     "ASCII Art Shadow: " + text,
		"thinkertoy": "ASCII Art Thinkertoy: " + text,
	}

	// Return the generated ASCII art based on the selected template
	return asciiMap[template]
}

// API Endpoint: /api/message - to generate ASCII Art
func messageHandler(w http.ResponseWriter, r *http.Request) {
	// Expecting POST with JSON or form data
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Get form values
	text := r.FormValue("text")
	template := r.FormValue("template")

	// If no text or template, return an error
	if text == "" || template == "" {
		http.Error(w, "Missing text or template", http.StatusBadRequest)
		return
	}

	// Generate ASCII Art based on template and text
	asciiArt := generateASCII(template, text)

	// Prepare response with ASCII Art
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: asciiArt})
}

// API Endpoint: /api/about - project info
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Message: "This is the ASCII art generation API! It allows you to convert text into various ASCII art templates."})
}

func main() {
	// Create a new HTTP request multiplexer (router)
	mux := http.NewServeMux()

	// Handle the /api/message endpoint for generating ASCII art
	mux.HandleFunc("/api/message", messageHandler)

	// Handle the /api/about endpoint for project information
	mux.HandleFunc("/api/about", aboutHandler)

	// Start the HTTP server with CORS middleware
	log.Println("Backend server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", enableCORS(mux)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
