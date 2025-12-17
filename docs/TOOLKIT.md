# GoShort URL Shortener — Toolkit Document

## 1) Title & Objective

**Tech:** Go (Golang) + Standard Library  
**Objective:** Build a minimal REST API for URL shortening to learn Go's backend capabilities and document learning with GenAI.

---

## 2) Quick Summary of the Technology

**Go:** Statically typed, compiled language from Google emphasizing simplicity, fast compilation, and native concurrency (goroutines).  
**Standard Library:** Built-in `net/http` package provides production-ready HTTP server—no framework needed.  
**Real-world usage:** Docker, Kubernetes, Terraform all built in Go. Used by Uber, Dropbox, Netflix for high-performance microservices.

---

## 3) System Requirements

- **OS:** Windows/macOS/Linux
- **Go:** 1.19+ (via official installer)
- **Tools:** Terminal, text editor (VS Code recommended)
- **Testing:** curl or Postman
- **Dependencies:** None (standard library only)

---

## 4) Installation & Setup

```bash
# Install Go (macOS)
brew install go

# Install Go (Linux)
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# Verify
go version

# Create project
mkdir go-url-shortener && cd go-url-shortener
go mod init goshort

# Create files
touch main.go handlers.go models.go utils.go

# Run server
go run .
# Open http://localhost:8080
```

---

## 5) Minimal Working Example

### Endpoints:
- `POST /shorten` → Create short URL `{"url": "https://example.com"}`
- `GET /:code` → Redirect to original URL
- `GET /stats/:code` → Get click statistics
- `GET /list` → List all URLs

### Expected Output:

**Start:** Server running on port 8080

**After POST /shorten:**
```json
{
  "short_code": "aBc123",
  "short_url": "http://localhost:8080/aBc123",
  "original_url": "https://example.com"
}
```

**GET /list shows:**
```json
[{
  "original_url": "https://example.com",
  "short_code": "aBc123",
  "clicks": 0,
  "created_at": "2025-12-16T20:25:20Z"
}]
```

---

## 6) AI Prompt Journal

See docs/ai-prompt-journal.md for full prompts, response summaries, and reflections.

## 7) Common Issues & Fixes

| Issue | Cause | Fix |
|-------|-------|-----|
| `expected 'package', found 'EOF'` | Incomplete .go file in directory | Remove test files: `rm test.go` |
| `address already in use` | Port 8080 occupied | Kill process: `lsof -i :8080` then `kill -9 <PID>`, or change port |
| `cannot find package` | Missing import statement | Add to imports: `import "net/url"` |
| JSON not parsing | Missing Content-Type header | Add `-H "Content-Type: application/json"` to curl |
| Redirect shows JSON error | Short code doesn't exist | Verify short_code from POST response before GET |
| Module not found | Not in project directory | Ensure in folder with go.mod: `cd goshort` |
| Code changes not reflected | Old process still running | Stop server (Ctrl+C) and restart with `go run .` |

---

## 8) References


- [Go Official Site](https://go.dev/) — Installation, tutorials, documentation
- [Go Tour](https://go.dev/tour/) — Interactive learning
- [Effective Go](https://go.dev/doc/effective_go) — Idiomatic Go patterns
- [Standard Library](https://pkg.go.dev/std) — Package reference

