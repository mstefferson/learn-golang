# Go Backend Mastery Course

**Duration:** ~8 weeks | 30 min/day, 5 days/week  
**Prerequisites:** Programming experience (Python or similar)  
**Capstone:** REST API with background worker queue

---

## Phase 1: Go Foundations (Week 1-2)

### Week 1 â€” Language Basics

| Day | Topic                       | Key Concepts                                                |
| --- | --------------------------- | ----------------------------------------------------------- |
| 1   | Environment & First Program | `go mod`, compilation model, `go run` vs `go build`         |
| 2   | Types & Variables           | Static typing, zero values, type inference, constants       |
| 3   | Control Flow                | `if`, `for` (no while), `switch`, `defer`                   |
| 4   | Functions                   | Multiple returns, named returns, variadic functions         |
| 5   | Error Handling              | Explicit errors, `error` interface, error checking patterns |

### Week 2 â€” Composite Types & Organization

| Day | Topic              | Key Concepts                                              |
| --- | ------------------ | --------------------------------------------------------- |
| 1   | Arrays & Slices    | Slice internals, capacity vs length, `append`, `copy`     |
| 2   | Maps & Structs     | Map operations, struct literals, embedding                |
| 3   | Pointers           | Address/dereference operators, pointer vs value semantics |
| 4   | Methods            | Value vs pointer receivers, method sets                   |
| 5   | Packages & Modules | Visibility rules, project structure, `go.mod`             |

---

## Phase 2: Interfaces & Dependency Injection (Week 3)

| Day | Topic                   | Key Concepts                                            |
| --- | ----------------------- | ------------------------------------------------------- |
| 1   | Interface Basics        | Implicit satisfaction, empty interface, type assertions |
| 2   | Standard Interfaces     | `io.Reader`, `io.Writer`, `fmt.Stringer`, `error`       |
| 3   | Constructor Injection   | Interface segregation, dependency inversion in Go       |
| 4   | Testing with Interfaces | Mocks, fakes, test doubles without frameworks           |
| 5   | **Mini-Project**        | Configurable logger with swappable backends             |

---

## Phase 3: Concurrency (Week 4-5)

### Week 4 â€” Primitives

| Day | Topic            | Key Concepts                                  |
| --- | ---------------- | --------------------------------------------- |
| 1   | Goroutines       | Spawning, Go scheduler model, stack growth    |
| 2   | Channels         | Unbuffered vs buffered, send/receive, closing |
| 3   | Select Statement | Multiplexing, timeouts, non-blocking ops      |
| 4   | Sync Package     | `Mutex`, `RWMutex`, `WaitGroup`, `Once`       |
| 5   | Context          | Cancellation, timeouts, value propagation     |

### Week 5 â€” Concurrency Patterns

| Day | Topic            | Key Concepts                              |
| --- | ---------------- | ----------------------------------------- |
| 1   | Worker Pools     | Fixed pool size, job distribution         |
| 2   | Fan-Out/Fan-In   | Parallel processing, result aggregation   |
| 3   | Pipelines        | Stage-based processing, channel chaining  |
| 4   | Rate Limiting    | Token bucket, semaphores, `time.Ticker`   |
| 5   | **Mini-Project** | Concurrent URL fetcher with rate limiting |

---

## Phase 4: Backend Development (Week 6)

| Day | Topic                | Key Concepts                                            |
| --- | -------------------- | ------------------------------------------------------- |
| 1   | HTTP Fundamentals    | `net/http`, handlers, `http.HandlerFunc`                |
| 2   | Routing & Middleware | Mux patterns, middleware chains, `http.Handler`         |
| 3   | JSON APIs            | Encoding/decoding, struct tags, validation              |
| 4   | Database Integration | `database/sql`, connection pooling, prepared statements |
| 5   | Production Concerns  | Graceful shutdown, health checks, signals               |

---

## Phase 5: Design Patterns in Go (Week 7)

| Day | Topic              | Key Concepts                                              |
| --- | ------------------ | --------------------------------------------------------- |
| 1   | Functional Options | Idiomatic configuration, variadic option funcs            |
| 2   | Repository Pattern | Clean architecture, layer separation                      |
| 3   | Factory & Builder  | Object construction patterns                              |
| 4   | Decorator Pattern  | Middleware as decorators, composable handlers             |
| 5   | Error Patterns     | Wrapping, `errors.Is`/`As`, sentinel errors, custom types |

---

## Phase 6: Capstone Project (Week 8)

### Project: Task Queue API

A production-style REST API with background job processing.

**Features:**

- RESTful task management endpoints
- Background worker pool for async processing
- PostgreSQL persistence with repository pattern
- Middleware stack (logging, auth, rate limiting)
- Graceful shutdown with in-flight job completion
- Health check and metrics endpoints

**Schedule:**

| Day | Focus                                                |
| --- | ---------------------------------------------------- |
| 1   | Architecture design, project scaffolding, DI setup   |
| 2   | Domain models, repository layer, database migrations |
| 3   | HTTP handlers, routing, middleware                   |
| 4   | Worker pool, job queue, graceful shutdown            |
| 5   | Testing, documentation, code review                  |

---

## Lesson Format

Each 30-minute session follows this structure:

| Duration | Activity                                    |
| -------- | ------------------------------------------- |
| 5 min    | Concept introduction with diagrams/examples |
| 20 min   | Hands-on coding exercises                   |
| 5 min    | Review, Q&A, preview of next lesson         |

---

## Resources

- [Go Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go Proverbs](https://go-proverbs.github.io/)

---

## Progress Tracker

- [ ] Week 1: Language Basics
- [ ] Week 2: Composite Types & Organization
- [ ] Week 3: Interfaces & Dependency Injection
- [ ] Week 4: Concurrency Primitives
- [ ] Week 5: Concurrency Patterns
- [ ] Week 6: Backend Development
- [ ] Week 7: Design Patterns
- [ ] Week 8: Capstone Project
