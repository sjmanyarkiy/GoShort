package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

// ShortenHandler creates a new short URL
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON request
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate URL
	if !IsValidURL(req.URL) {
		sendError(w, "Invalid URL provided", http.StatusBadRequest)
		return
	}

	// Generate unique short code
	shortCode := GenerateShortCode()
	for urlStore[shortCode] != nil {
		shortCode = GenerateShortCode()
	}

	// Create mapping
	mapping := &URLMapping{
		OriginalURL: req.URL,
		ShortCode:   shortCode,
		Clicks:      0,
		CreatedAt:   time.Now(),
	}

	// Store mapping
	urlStore[shortCode] = mapping

	log.Printf("Created short URL: %s -> %s", shortCode, req.URL)

	// Send response
	response := ShortenResponse{
		ShortCode:   shortCode,
		ShortURL:    "http://localhost:8080/" + shortCode,
		OriginalURL: req.URL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RedirectHandler redirects short code to original URL
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Extract short code from path
	shortCode := strings.TrimPrefix(r.URL.Path, "/")

	// Handle special routes
	if shortCode == "" || shortCode == "shorten" || 
	   strings.HasPrefix(shortCode, "stats/") || shortCode == "list" {
		http.NotFound(w, r)
		return
	}

	// Look up mapping
	mapping := urlStore[shortCode]
	if mapping == nil {
		sendError(w, "Short code not found", http.StatusNotFound)
		return
	}

	// Increment click counter
	mapping.Clicks++

	log.Printf("Redirecting %s -> %s (clicks: %d)", shortCode, mapping.OriginalURL, mapping.Clicks)

	// Redirect to original URL
	http.Redirect(w, r, mapping.OriginalURL, http.StatusMovedPermanently)
}

// StatsHandler returns statistics for a short code
func StatsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract short code from path
	shortCode := strings.TrimPrefix(r.URL.Path, "/stats/")

	// Look up mapping
	mapping := urlStore[shortCode]
	if mapping == nil {
		sendError(w, "Short code not found", http.StatusNotFound)
		return
	}

	// Send stats response
	response := StatsResponse{
		OriginalURL: mapping.OriginalURL,
		ShortCode:   mapping.ShortCode,
		Clicks:      mapping.Clicks,
		CreatedAt:   mapping.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ListHandler returns all stored URLs
func ListHandler(w http.ResponseWriter, r *http.Request) {
	urls := make([]StatsResponse, 0, len(urlStore))

	for _, mapping := range urlStore {
		urls = append(urls, StatsResponse{
			OriginalURL: mapping.OriginalURL,
			ShortCode:   mapping.ShortCode,
			Clicks:      mapping.Clicks,
			CreatedAt:   mapping.CreatedAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)
}

// sendError sends a JSON error response
func sendError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}