package main

import (
	"fmt"
	"net/http"
)

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Start of our ascii art")
}

func main() {
	http.HandleFunc("/ascii", HandleFunc)
	http.ListenAndServe(":8080", nil)
	
}