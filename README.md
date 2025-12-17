# GoShort - URL Shortener API

A simple and efficient URL shortening service built with Go.

## 🚀 Quick Start

### Prerequisites
- Go 1.19 or higher installed
- Basic understanding of REST APIs
- Editor: VS Code

### Installation

```bash
# Clone the project
git clone https://github.com/sjmanyarkiy/GoShort.git
cd Goshort

# Run the server
go run .
```

The server will start on `http://localhost:8080`

## 📡 API Endpoints

### 1. Shorten a URL

**POST** `/shorten`

Create a shortened URL from a long URL.

**Request:**
```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://www.example.com/very/long/url"}'
```

**Response:**
```json
{
  "short_code": "aB3xYz",
  "short_url": "http://localhost:8080/aB3xYz",
  "original_url": "https://www.example.com/very/long/url"
}
```

### 2. Redirect to Original URL

**GET** `/:shortCode`

Redirects to the original URL and increments the click counter.

**Example:**
```bash
curl -L http://localhost:8080/aB3xYz
```

Or simply visit in browser: `http://localhost:8080/aB3xYz`

### 3. Get Statistics

**GET** `/stats/:shortCode`

Returns statistics for a shortened URL.

**Request:**
```bash
curl http://localhost:8080/stats/aB3xYz
```

**Response:**
```json
{
  "original_url": "https://www.example.com/very/long/url",
  "short_code": "aB3xYz",
  "clicks": 5,
  "created_at": "2024-01-15T10:30:00Z"
}
```

### 4. List All URLs

**GET** `/list`

Returns all shortened URLs with their statistics.

**Request:**
```bash
curl http://localhost:8080/list
```

**Response:**
```json
[
  {
    "original_url": "https://www.example.com",
    "short_code": "aB3xYz",
    "clicks": 5,
    "created_at": "2024-01-15T10:30:00Z"
  },
  {
    "original_url": "https://www.google.com",
    "short_code": "xY9zAb",
    "clicks": 12,
    "created_at": "2024-01-15T11:45:00Z"
  }
]
```

## 🧪 Testing Examples

### Using curl

```bash
# 1. Create a short URL
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://github.com"}'

# Expected output: {"short_code":"aB3xYz","short_url":"http://localhost:8080/aB3xYz","original_url":"https://github.com"}

# 2. Test the redirect (use -L to follow redirects)
curl -L http://localhost:8080/aB3xYz

# 3. Check statistics
curl http://localhost:8080/stats/aB3xYz

# 4. List all URLs
curl http://localhost:8080/list
```

### Using Postman

1. **Create Short URL:**
   - Method: POST
   - URL: `http://localhost:8080/shorten`
   - Headers: `Content-Type: application/json`
   - Body (raw JSON):
     ```json
     {
       "url": "https://www.example.com"
     }
     ```

2. **Get Stats:**
   - Method: GET
   - URL: `http://localhost:8080/stats/{shortCode}`

3. **List All:**
   - Method: GET
   - URL: `http://localhost:8080/list`

## 🏗️ Project Structure

```
go-url-shortener/
├── main.go          # HTTP server setup and routing
├── handlers.go      # Request handlers for each endpoint
├── models.go        # Data structures (URLMapping, requests, responses)
├── utils.go         # Utility functions (code generation, validation)
├── go.mod           # Go module file
└── README.md        # This file
└── TOOLKIT.md       # AI usage and tooling notes
```

## 🔧 How It Works

1. **URL Shortening**: Generates a random 6-character alphanumeric code
2. **Storage**: Uses in-memory map for storing URL mappings
3. **Redirects**: HTTP 301 permanent redirect to original URL
4. **Statistics**: Tracks clicks and creation timestamp for each URL

## ⚠️ Limitations

- **In-Memory Storage**: Data is lost when server restarts
- **No Persistence**: URLs are not saved to database
- **No Authentication**: All endpoints are public
- **Collision Handling**: Simple regeneration if code exists

## 🚀 Future Enhancements

- [ ] Database persistence (PostgreSQL/MongoDB)
- [ ] Custom short codes
- [ ] Expiration dates for URLs
- [ ] User authentication
- [ ] Rate limiting
- [ ] Analytics dashboard
- [ ] API key management

## 📝 Error Handling

The API returns appropriate HTTP status codes:

- `200 OK` - Successful request
- `301 Moved Permanently` - Successful redirect
- `400 Bad Request` - Invalid URL or malformed JSON
- `404 Not Found` - Short code doesn't exist
- `405 Method Not Allowed` - Wrong HTTP method
- `500 Internal Server Error` - Server error

## 👨‍💻 Development

```bash
# Run the server
go run .

# Build executable
go build -o goshort

# Run the executable
./goshort
```

## References

- Go Documentation
- Moringa School AI Curriculum

## 📄 License

This project is open source and available under the MIT License.

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!

---
## Author
Built by Sandra Manyarkiy with ❤️ using Go for the "AI Artificial Intelligence" capstone project