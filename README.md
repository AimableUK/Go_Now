# Go_Now

Welcome to my **Go (Golang)** learning repository! This project serves as a dedicated workspace for exploring the Go programming language, mastering its idiomatic patterns, and building robust, high-performance backend services.

---

## 🎯 Project Objectives

- **Core Proficiency:** Mastering Go syntax, concurrency models, and standard library best practices.
- **Backend Development:** Architecting RESTful APIs and microservices.
- **Ecosystem Exploration:** Hands-on experience with modern frameworks like [Gin](https://gin-gonic.com/) and standard library-first development.

## 🗺️ Learning Roadmap

### Phase 01 - Go Basics (Completed)

- [x] **Basics & Syntax**
  - Hello World
  - Variables & Constants
  - Printing & Formatting
  - Arrays & Slices
  - Control Flow (If/Else, Switch)
  - Loops (`for`, `range`)
  - Functions & Recursion

### Phase 02 - Mastering Go

- [ ] **Go by Example**
  - **Data & Memory**
    - [ ] Pointers
    - [ ] Structs
    - [ ] Struct Embedding
    - [ ] Enums
    - [ ] Map
    - [ ] Methods
    - [ ] Interfaces
    - [ ] Generics
  - **Advanced Control Flow**
    - [ ] Closures
    - [ ] Variadic Functions
    - [ ] Multiple Return Values
    - [ ] Recursion
    - [ ] Defer/Panic/Recover
    - [ ] Range over Iterators
  - **Error Handling**
    - [ ] Errors
    - [ ] Custom Errors
    - [ ] Error Wrapping (`%w`)
    - [ ] `errors.Is` / `errors.As`
    - [ ] Sentinel Errors
  - **Concurrency (The Go Way)**
    - [ ] Goroutines
    - [ ] Channels
    - [ ] Channel Buffering
    - [ ] Channel Directions
    - [ ] Channel Synchronization
    - [ ] Closing/Range over Channels
    - [ ] Select
    - [ ] Non-Blocking Channel Operations
    - [ ] Timeouts
    - [ ] Timers
    - [ ] Tickers
    - [ ] WaitGroups
    - [ ] Mutexes
    - [ ] Atomic Counters
    - [ ] Stateful Goroutines
    - [ ] Worker Pools
    - [ ] Rate Limiting
  - **Data Handling & I/O**
    - [ ] JSON
    - [ ] XML
    - [ ] String Functions & Formatting
    - [ ] Text Templates
    - [ ] Regular Expressions
    - [ ] Sorting (& Sorting by Functions)
    - [ ] Time/Epoch
    - [ ] Time Formatting/Parsing
    - [ ] Number Parsing
    - [ ] URL Parsing
    - [ ] SHA256 Hashes
    - [ ] Base64 Encoding
  - **File & OS Operations**
    - [ ] Reading/Writing Files
    - [ ] File Paths
    - [ ] Directories
    - [ ] Temporary Files
    - [ ] Embed Directive
    - [ ] Environment Variables
    - [ ] Command-Line Arguments/Flags/Subcommands
    - [ ] Signals
    - [ ] Exit
    - [ ] Spawning/Exec'ing Processes
  - **Systems & Networking**
    - [ ] HTTP Client
    - [ ] HTTP Server
    - [ ] TCP Server
    - [ ] Context
  - **Testing & Tooling**
    - [ ] Testing
    - [ ] Benchmarking
    - [ ] Logging (`log/slog`)
    - [ ] `gofmt`
    - [ ] `go vet`
    - [ ] `golangci-lint`
    - [ ] `pprof` Profiling

- [ ] **Go Modules & Project Hygiene**
  - [ ] `go.mod` / `go.sum` Management
  - [ ] Semantic Versioning
  - [ ] Go Workspaces (`go.work`)

- [ ] **RESTful API with Go & Gin**
  - [ ] Gin Framework Setup
  - [ ] Routing & Context
  - [ ] JSON Handling
  - [ ] Middleware

- [ ] **Building a Task Manager REST API**
  - [ ] Project Structure
  - [ ] CRUD Operations
  - [ ] Error Handling
  - [ ] Request Validation

- [ ] **Postman Documentation & Integration**
  - [ ] API Collections
  - [ ] Environment Variables
  - [ ] Automated Testing with Postman

- [ ] **MongoDB Fundamentals & Driver Tutorial**
  - [ ] MongoDB Setup
  - [ ] Connecting Go to MongoDB
  - [ ] Basic CRUD Queries
  - [ ] Document Modeling

- [ ] **SQL Databases in Go**
  - [ ] `database/sql` Fundamentals
  - [ ] `sqlc` or GORM
  - [ ] Migrations (`golang-migrate`)

- [ ] **Authentication & Authorization in Go**
  - [ ] JWT (JSON Web Tokens)
  - [ ] Password Hashing (bcrypt)
  - [ ] Middleware for Auth

- [ ] **Clean Architecture Principles (Uncle Bob)**
  - [ ] Entities, Use Cases, Adapters
  - [ ] Dependency Rule
  - [ ] Dependency Injection

- [ ] **Implementing Clean Architecture in Go**
  - [ ] Refactoring to Clean Architecture
  - [ ] Separating Concerns

- [ ] **Testing & Quality Assurance (Testify & Mockery)**
  - [ ] Unit Testing
  - [ ] Testify Assertions
  - [ ] Mocking with Mockery
  - [ ] Testing REST API Endpoints

- [ ] **Observability & Deployment**
  - [ ] Health Checks
  - [ ] OpenTelemetry Basics
  - [ ] Docker Multi-Stage Builds for Go Binaries
  - [ ] Basic CI with GitHub Actions

- [ ] **gRPC in Go**
  - [ ] Protocol Buffers
  - [ ] Service Definitions
  - [ ] REST + gRPC Coexistence

---

## 📂 Project Structure

```text
Go_Now/
├── main.go      # Entry point of the application
├── go.mod       # Dependency management
└── README.md    # Documentation
```

## Getting Started

1. **Prerequisites:**
   Ensure you have [Go installed](https://go.dev/doc/install) (version 1.20+ recommended).
2. **Run the Project:**
   ```bash
   go run main.go
   ```

---

## 🤝 Contributions

This is a personal learning repository. However, feel free to open an issue if you have suggestions or find improvements for the code snippets provided.

_Happy Coding!_ 🐹
