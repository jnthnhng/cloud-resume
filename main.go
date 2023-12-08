package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve src files from the "src" directory
	fs := http.FileServer(http.Dir("src"))
	http.Handle("/src/", http.StripPrefix("/src/", fs))

	// Serve the HTML file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "src/index.html")
	})

	// Start the HTTP server on port 8080
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
