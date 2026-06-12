# Go Project: Minimal REST API Server (Notes Service)

## Project Overview

Build a small REST API server in Go that manages simple notes.

This project introduces:

- ✅ HTTP server development with `net/http`
- ✅ JSON encoding and decoding
- ✅ Routing basics
- ✅ Middleware pattern
- ✅ Concurrency safety with `sync.RWMutex`
- ✅ Clean project structure
- ✅ Context awareness
- ✅ Proper HTTP status codes

This is a foundational backend engineering project in Go.

---

# What You Will Build

A simple Notes API with in-memory storage.

## Endpoints

### Create a Note

```
POST /notes
```

Request body:

```json
{
  "title": "Learn Go",
  "content": "Practice building REST APIs"
}
```

Response:

```json
{
  "id": 1,
  "title": "Learn Go",
  "content": "Practice building REST APIs"
}
```

---

### Get All Notes

```
GET /notes
```

Response:

```json
[
  {
    "id": 1,
    "title": "Learn Go",
    "content": "Practice building REST APIs"
  }
]
```

---

### Get Single Note

```
GET /notes/{id}
```

---

### Delete Note

```
DELETE /notes/{id}
```

---

# Learning Objectives

After completing this project, you should understand:

- How Go’s `net/http` server works
- How handlers process requests
- How to encode/decode JSON
- How to manage shared state safely
- How middleware wraps handlers
- How to structure small backend services

---

# Architecture Overview

```
HTTP Request
      ↓
Router (ServeMux)
      ↓
Middleware (Logging)
      ↓
Handler
      ↓
In-memory Store (protected by Mutex)
      ↓
JSON Response
```

---

# Suggested Project Structure

```
notes-api/
│
├── go.mod
├── main.go
├── handlers.go
├── store.go
├── middleware.go
└── models.go
```

For a small project, you may also implement everything inside `main.go`.

---

# Core Components

---

## 1. Note Model

```go
type Note struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
}
```

You will learn about:

- Struct tags
- JSON serialization
- Data modeling

---

## 2. In-Memory Store

```go
type Store struct {
    mu    sync.RWMutex
    notes map[int]Note
    nextID int
}
```

Why `RWMutex`?

- Multiple reads allowed simultaneously
- Writes require exclusive lock
- Prevents race conditions

This introduces real-world concurrency safety.

---

## 3. HTTP Handlers

Each handler has this signature:

```go
func(w http.ResponseWriter, r *http.Request)
```

Responsibilities:

- Validate HTTP method
- Parse request body (if needed)
- Access store safely
- Return JSON
- Set correct HTTP status codes

---

## 4. JSON Handling

Decode request:

```go
json.NewDecoder(r.Body).Decode(&note)
```

Encode response:

```go
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(data)
```

---

## 5. Middleware Pattern

Example: Logging middleware

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
```

This demonstrates:

- Handler wrapping
- Chainable architecture
- Cross-cutting concerns

Middleware is widely used in production Go servers.

---

# Implementation Roadmap

---

## Step 1 – Initialize Project

```
mkdir notes-api
cd notes-api
go mod init notes-api
```

---

## Step 2 – Implement Basic Server

Start with:

```go
http.ListenAndServe(":8080", nil)
```

Confirm server runs.

---

## Step 3 – Add Routing

Use `http.NewServeMux()`:

```go
mux := http.NewServeMux()
mux.HandleFunc("/notes", handleNotes)
mux.HandleFunc("/notes/", handleNoteByID)
```

---

## Step 4 – Add In-Memory Store

Implement:

- Create note
- Get all notes
- Get note by ID
- Delete note

Protect access with `RWMutex`.

---

## Step 5 – Add Middleware

Wrap your mux:

```go
loggedMux := loggingMiddleware(mux)
http.ListenAndServe(":8080", loggedMux)
```

---

## Step 6 – Test with curl

Create note:

```
curl -X POST http://localhost:8080/notes \
-H "Content-Type: application/json" \
-d '{"title":"Test","content":"Hello"}'
```

Get notes:

```
curl http://localhost:8080/notes
```

Delete note:

```
curl -X DELETE http://localhost:8080/notes/1
```

---

# HTTP Status Codes to Use

| Situation          | Code |
| ------------------ | ---- |
| Success (GET)      | 200  |
| Created (POST)     | 201  |
| Not Found          | 404  |
| Bad Request        | 400  |
| Method Not Allowed | 405  |

Learning proper HTTP semantics is important for backend work.

---

# Possible Enhancements

Once basic version works, try:

- ✅ Add request validation
- ✅ Add update endpoint (`PUT /notes/{id}`)
- ✅ Add graceful shutdown with `context`
- ✅ Add request timeout
- ✅ Add structured logging
- ✅ Add environment-based configuration
- ✅ Add simple unit tests
- ✅ Add persistence (JSON file storage)

---

# Advanced Version (Optional Challenge)

Add:

- Context cancellation
- Request ID middleware
- Rate limiting
- Metrics endpoint
- Dockerfile
- Benchmarking

This moves you toward production-level services.

---

# What This Project Represents

You are transitioning from:

- CLI tools
- Concurrency exercises
- Networking utilities

To:

- Service-oriented backend development
- Stateful systems
- API design
- Production-ready architecture patterns

This is the foundation of real-world Go backend engineering.

---

# After Completing This

Strong next steps:

- Reverse proxy server
- Authentication middleware
- Background job worker
- Persistent storage with SQLite
- gRPC service
- Microservice-style architecture

---

# Final Goal

By finishing this project, you should feel comfortable building:

- Small internal APIs
- Backend services
- Microservices
- Infrastructure components

You are now moving from "learning Go syntax"
to "engineering Go systems."
