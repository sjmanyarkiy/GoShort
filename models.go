package main

import "time"

// URLMapping represents a shortened URL with its metadata
type URLMapping struct {
	OriginalURL string    `json:"original_url"`
	ShortCode   string    `json:"short_code"`
	Clicks      int       `json:"clicks"`
	CreatedAt   time.Time `json:"created_at"`
}

// ShortenRequest is the JSON payload for creating a short URL
type ShortenRequest struct {
	URL string `json:"url"`
}

// ShortenResponse is returned after successfully creating a short URL
type ShortenResponse struct {
	ShortCode   string `json:"short_code"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

// StatsResponse contains statistics for a shortened URL
type StatsResponse struct {
	OriginalURL string    `json:"original_url"`
	ShortCode   string    `json:"short_code"`
	Clicks      int       `json:"clicks"`
	CreatedAt   time.Time `json:"created_at"`
}

// ErrorResponse is a standard error message format
type ErrorResponse struct {
	Error string `json:"error"`
}