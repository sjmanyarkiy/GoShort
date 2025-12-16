# GoShort – URL Shortener API

A minimal REST API built with Go to shorten URLs, redirect users, and track basic usage statistics.

## Features

* Shorten long URLs into compact codes
* Redirect short URLs to original destinations
* Track click counts and creation time
* Simple JSON-based REST API
* In-memory storage for fast access

## Requirements

* Go 1.19+
* Git
* Editor: VS Code (recommended)

## Setup Instructions

```bash
# Clone the repository
git clone https://github.com/sjmanyarkiy/GoShort.git
cd GoShort

# Run the application
go run .
```

Server runs at: `http://localhost:8080`

## API Endpoints

### ➕ POST /shorten

Create a new shortened URL.

**Request**

```json
{
  "url": "https://www.example.com/very/long/url"
}
```

**Response — 201 Created**

```json
{
  "short_code": "aB3xYz",
  "short_url": "http://localhost:8080/aB3xYz",
  "original_url": "https://www.example.com/very/long/url"
}
```

---

### 🔁 GET /:shortCode

Redirect to the original URL and increment the click counter.

**Example**

```bash
curl -L http://localhost:8080/aB3xYz
```

---

### 📊 GET /stats/:shortCode

Retrieve statistics for a shortened URL.

**Response**

```json
{
  "original_url": "https://www.example.com/very/long/url",
  "short_code": "aB3xYz",
  "clicks": 5,
  "created_at": "2024-01-15T10:30:00Z"
}
```

---

### 📄 GET /list

Retrieve all stored URLs and their statistics.

**Response**

```json
[
  {
    "original_url": "https://www.example.com",
    "short_code": "aB3xYz",
    "clicks": 5,
    "created_at": "2024-01-15T10:30:00Z"
  }
]
```

## Common Issues

| Error                    | Solution                                       |
| ------------------------ | ---------------------------------------------- |
| Address already in use   | Stop the running server or change the port     |
| Invalid JSON body        | Ensure `Content-Type: application/json` is set |
| 404 short code not found | Verify the short code exists                   |
| Data lost on restart     | Expected behavior (in-memory storage)          |

## Project Structure

```
GoShort/
├── main.go        # Application entry point and routing
├── handlers.go    # HTTP handlers
├── models.go      # Data models and DTOs
├── utils.go       # Helpers (validation, code generation)
├── go.mod         # Go module definition
├── README.md      # Project documentation
└── TOOLKIT.md     # AI usage and tooling notes
```

## Notes & Limitations

* Uses in-memory storage (no persistence)
* No authentication or rate limiting
* Not intended for production use without extensions

## References

* Go Documentation
* net/http Package
* REST API Design Guidelines

## AI Prompt Journal — GoShort URL Shortener API

**Source → Target:** JavaScript/React → Go (net/http)

This journal documents AI prompts, summarized responses, and reflections following a structured 4‑step learning flow while building GoShort.

---

## Step 1: Conceptual Understanding

### Prompt (used as-is)

I'm proficient in JavaScript/React and want to learn Go for building backend REST APIs.
Before writing code:

1. What are the key philosophical differences between JavaScript and Go?
2. What problems was Go designed to solve?
3. What mental models should I adjust coming from JavaScript?
4. What are common misconceptions JavaScript developers have about Go?

### Response Summary (concise)

* **Philosophy:** JavaScript prioritizes flexibility and rapid iteration; Go prioritizes simplicity, readability, and predictability.
* **Problems solved:** Fast compilation, easy concurrency, low operational complexity, and maintainable backend services.
* **Mental shifts:** Explicit error handling, static typing, minimal abstractions, goroutines instead of event-loop thinking.
* **Misconceptions:** “Go is too simple,” “Go can’t scale,” “No generics means weak design.” Go favors clarity over cleverness.

### Reflection

Coming from JavaScript, I need to write more explicit code and handle errors intentionally. Go trades expressiveness for clarity, which should improve long-term maintainability and debugging.

---

## Step 2: Step-by-Step Breakdown (Concept Focus)

### Concept Chosen: Concurrency & HTTP Servers in Go

### Prompt

I want to understand concurrency in Go for REST APIs. Please explain:

1. How goroutines and channels work conceptually
2. How this compares to JavaScript’s event loop and promises
3. When to use mutexes vs channels
4. Best practices for shared state in net/http servers
   Focus on concepts, not complex code.

### Response Summary

* **Goroutines:** Lightweight threads managed by the Go runtime.
* **Comparison:** JS uses a single-threaded event loop; Go uses multi-threaded concurrency with CSP principles.
* **Mutex vs channels:** Mutexes protect shared memory; channels coordinate communication.
* **Best practices:** Minimize shared state, guard maps with mutexes, avoid blocking handlers.

### Reflection

Go’s concurrency model feels more explicit than JavaScript. I must be careful with shared state in handlers, especially when tracking URL statistics concurrently.

---

## Step 3: Guided Implementation

### Prompt

I want to build a simple URL shortener in Go.
Guide me to implement:

* POST /shorten
* GET /:shortCode redirect
* In-memory storage with concurrency safety
  Explain how each part differs from a JavaScript/Express implementation.

### Response Summary

* Use structs for request/response models.
* Use map[string]URLMapping protected by sync.Mutex.
* net/http handlers replace Express middleware.
* Manual routing and JSON encoding via encoding/json.

### Reflection

Handlers feel lower-level than Express, but also clearer. Managing state explicitly with mutexes forced me to think about concurrency earlier than I would in JavaScript.

---

## Step 4: Understanding Verification

### Prompt Template (after implementation)

I've implemented this Go URL shortener:

``` 
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
```

Please:

1. Verify Go best practices
2. Suggest improvements (errors, structure, concurrency)
3. Recommend next learning steps
4. Identify JavaScript habits in my Go code

### Verification Checklist

* Avoid panic or log.Fatal inside handlers
* Validate input URLs before storing
* Lock shared maps correctly
* Generate short codes server-side
* Separate handlers, models, and utilities

### Next Topics

* Testing with net/http/httptest
* Structured logging
* Database persistence
* Rate limiting and middleware
* Graceful shutdowns

---

## Log Table

| Date       | Step | Prompt Focus      | Response Summary           | Reflection / Change                 |
| ---------- | ---- | ----------------- | -------------------------- | ----------------------------------- |
| 2025-01-10 | 1    | Go vs JS concepts | Philosophy & mental models | Became explicit with error handling |
| 2025-01-12 | 2    | Concurrency       | Goroutines & mutexes       | Added mutex to URL map              |
| 2025-01-14 | 3    | Core endpoints    | POST/GET handlers          | Simplified handler logic            |
| 2025-01-15 | 4    | Verification      | Improvements & structure   | Cleaned up project layout           |

---

## Author

Built by **Sandra Manyarkiy** as a learning-focused capstone project using Go and AI-assisted guidance.
