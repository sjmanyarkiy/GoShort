# GoShort - AI Toolkit & Prompt Journal

**Student Name:** Sandra Manyarkiy  
**Project:** GoShort - URL Shortener API  
**Technology:** Go (Golang)  
**Completion Date:** December 16, 2025

---

## 📚 AI Tools Used

### Primary Tool: Claude AI (Anthropic)
- **Purpose:** Code generation, debugging, documentation
- **Platform:** claude.ai
- **Model:** Claude Sonnet 4.5

---

## 🤖 Prompt Strategy & Journal

### Day 1: Project Setup

**Prompt 1: Environment Setup**
```
I want to create an application using Go. I have a project plan for a URL 
shortener. Can you help me set up Go and create the project structure?
```

**AI Response Summary:**
- Provided installation instructions for Go
- Created project structure with 4 files
- Generated complete boilerplate code

**What I Learned:**
- How to initialize a Go module with `go mod init`
- Basic Go project structure
- Package management in Go

---

### Day 2: Core Development

**Prompt 2: Understanding the Code**
```
Explain how the URL shortening logic works in the handlers.go file
```

**Key Concepts Learned:**
- HTTP handlers in Go
- JSON encoding/decoding
- Map-based storage
- Random string generation

**Prompt 3: Debugging**
```
When I run "go run ." I get "expected 'package', found 'EOF'" error
```

**Solution Applied:**
- Removed incomplete test.go file
- Learned about Go's compilation process
- Fixed project directory structure

---

### Day 3: Testing & Refinement

**Prompt 4: Testing Strategy**
```
Create a comprehensive testing checklist for my URL shortener API
```

**Testing Approach:**
- Used curl for API testing
- Tested all endpoints (POST, GET, redirects)
- Validated error handling
- Checked edge cases

---

## 🔧 Technical Challenges & Solutions

### Challenge 1: Understanding Go Syntax
**Problem:** Coming from JavaScript background, Go syntax was different  
**Solution:** AI explained key differences:
- Statically typed vs dynamically typed
- Explicit error handling (no try-catch)
- Pointers and memory management

### Challenge 2: HTTP Server Setup
**Problem:** Didn't know how to set up HTTP server in Go  
**Solution:** AI provided complete server code with:
- Route handlers
- Middleware concepts
- Request/response handling

### Challenge 3: Data Storage
**Problem:** How to store URLs without a database  
**Solution:** Used in-memory map structure
- `var urlStore = make(map[string]*URLMapping)`
- Fast lookups with O(1) complexity

---

## 💡 Key Learnings

### Go-Specific Concepts

1. **Structs vs Classes**
   - Go uses structs instead of classes
   - Methods can be attached to structs
   - JSON tags for serialization

2. **Error Handling**
   - Explicit error returns
   - No exceptions in Go
   - Pattern: `if err != nil { return err }`

3. **HTTP Package**
   - `net/http` is built-in
   - Simple handler functions
   - Minimal framework needed

4. **Concurrency**
   - Go's built-in HTTP server handles concurrent requests
   - No explicit thread management needed

### Software Development Principles

1. **API Design**
   - RESTful endpoint structure
   - Proper HTTP status codes
   - JSON request/response format

2. **Code Organization**
   - Separation of concerns (models, handlers, utils)
   - Single responsibility principle
   - Clean file structure

3. **Testing**
   - Test all happy paths
   - Test error cases
   - Document test results

---

## 🎯 Effective AI Prompting Techniques

### What Worked Well

1. **Be Specific**
   - ✅ "Create a POST endpoint in Go that accepts JSON"
   - ❌ "Help me with Go"

2. **Provide Context**
   - Share error messages verbatim
   - Mention your background (JavaScript → Go)
   - Explain the end goal

3. **Iterative Refinement**
   - Start with basic implementation
   - Ask for improvements
   - Request explanations for unclear parts

4. **Ask for Explanations**
   - Don't just copy code
   - Ask "Why does this work?"
   - Request analogies to familiar concepts

### Prompts That Generated Best Results

1. "Explain [concept] as if I'm coming from JavaScript"
2. "Show me the idiomatic Go way to [task]"
3. "What are common mistakes beginners make with [feature]?"
4. "Create a complete example with error handling"

---

## 📊 Project Statistics

- **Total Time:** ~4 hours
- **Lines of Code:** ~250
- **Files Created:** 5 (4 Go files + README)
- **API Endpoints:** 4
- **AI Interactions:** ~10 major prompts

---

## 🚀 Future Improvements (If I Had More Time)

1. **Database Integration**
   - Add PostgreSQL or MongoDB
   - Persist data across restarts

2. **Custom Short Codes**
   - Let users choose their own codes
   - Validate availability

3. **Analytics Dashboard**
   - Web UI for viewing stats
   - Charts for click trends

4. **Authentication**
   - User accounts
   - API key management

5. **Advanced Features**
   - URL expiration
   - QR code generation
   - Rate limiting

---

## 🎓 Reflection

### What Went Well
- AI helped me learn Go syntax quickly
- Project completed faster than expected
- Code is clean and well-organized
- All features work as intended

### What Was Challenging
- Understanding Go's error handling pattern
- Different mindset from JavaScript
- Learning when to use pointers

### How AI Enhanced My Learning
- Instant feedback on errors
- Code examples with explanations
- Best practices and idioms
- Saved hours of documentation reading

### Skills Gained
- Go programming fundamentals
- REST API design
- HTTP server implementation
- JSON handling
- Error management
- Testing methodologies

---

## 📝 Code Quality Checklist

- [x] Code is readable and well-formatted
- [x] Functions have clear purposes
- [x] Error handling implemented
- [x] All endpoints tested
- [x] Documentation complete
- [x] No hardcoded values (except port)
- [x] Proper HTTP status codes used
- [x] JSON responses formatted correctly

---

## 🙏 Acknowledgments

- **Claude AI** for code generation and explanations
- **Go Documentation** for reference
- **Project Plan** for structured approach

---

**Final Note:** This project demonstrates that AI tools are powerful learning 
accelerators when used thoughtfully. The key is asking good questions, 
understanding the responses, and iterating based on results.