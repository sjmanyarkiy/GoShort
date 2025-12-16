package main

import (
	"log"
	"net/http"
	"strings"
)

// Global in-memory storage for URLs
var urlStore = make(map[string]*URLMapping)

func main() {
	// Setup routes
	http.HandleFunc("/shorten", ShortenHandler)
	http.HandleFunc("/list", ListHandler)
	http.HandleFunc("/stats/", StatsHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Route to appropriate handler
		path := r.URL.Path
		
		if path == "/" {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message":"GoShort URL Shortener API","endpoints":["/shorten","/list","/stats/:code","/:code"]}`))
			return
		}
		
		if strings.HasPrefix(path, "/stats/") {
			StatsHandler(w, r)
			return
		}
		
		// Default to redirect handler
		RedirectHandler(w, r)
	})

	// Start server
	port := ":8080"
	log.Printf("🚀 GoShort server starting on http://localhost%s", port)
	log.Println("📍 Endpoints:")
	log.Println("   POST /shorten      - Create short URL")
	log.Println("   GET  /:code        - Redirect to original URL")
	log.Println("   GET  /stats/:code  - Get URL statistics")
	log.Println("   GET  /list         - List all URLs")
	
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}