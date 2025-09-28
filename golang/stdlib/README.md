## Phase 1: Foundations of Go Networking

Goal: Understand how Go abstracts sockets and networking.

1. **`net` package (low-level networking)**

   * Learn about `net.Listener`, `net.Conn`
   * Write:

     * TCP echo server and client
     * UDP message sender/receiver
     * Simple port scanner
   * Concepts: blocking I/O, deadlines, connection lifecycle

2. **`bufio` & `io`**

   * Learn buffered I/O on top of raw connections.
   * Implement your own line-based protocol reader/writer.
   * Practice streaming large files over TCP.

3. **`context`**

   * Explore cancellation, timeouts, and deadlines.
   * Add graceful shutdown to your TCP server.

---

## Phase 2: High-Level HTTP

Goal: Master Go’s `net/http` stack.

1. **`net/http` basics**

   * Start with `http.ListenAndServe`
   * Build a toy REST API
   * Learn how handlers and middleware chaining works

2. **Custom `http.Server`**

   * Control server lifecycle
   * Add timeouts (`ReadTimeout`, `WriteTimeout`)
   * Graceful shutdown with `Server.Shutdown(ctx)`

3. **HTTP client (`http.Client`)**

   * Timeouts, retries, connection pooling
   * Follow redirects manually
   * Try chunked transfer decoding

4. **Multiplexing and routing**

   * Implement a mini-router using `ServeMux`
   * Compare with third-party routers (chi, gorilla/mux)

---

## Phase 3: Concurrency & Robustness

Goal: Apply Go concurrency patterns to networking.

1. **Goroutines & channels**

   * Worker pool handling multiple TCP clients
   * Broadcast messages to all connected clients

2. **`sync` and `atomic`**

   * Safe connection counters
   * Shared resource access in servers

3. **`errgroup` (x/sync)**

   * Coordinate multiple goroutines
   * Manage background tasks in servers

4. **Graceful shutdown**

   * OS signals (`os/signal`, `syscall`)
   * Canceling goroutines with `context`

---

## Phase 4: Advanced HTTP

Goal: Understand what makes `net/http` production-grade.

1. **HTTP/2**

   * Enable HTTP/2 with `http2` package (x/net/http2)
   * Compare request multiplexing vs HTTP/1.1
   * Play with server push (deprecated but educational)

2. **TLS (`crypto/tls`)**

   * Load certificates, start HTTPS server
   * Explore mutual TLS (client certificates)

3. **Reverse proxying**

   * Explore `httputil.ReverseProxy`
   * Write a simple load balancer (round-robin across multiple backends)

4. **Middlewares**

   * Logging, authentication, rate limiting
   * Explore request context (`r.Context()`)

---

## Phase 5: Best Practices & Tooling

Goal: Develop production-ready habits.

1. **Logging**

   * Structured logging (`log/slog`)
   * Middleware logging for requests

2. **Profiling & performance**

   * `net/http/pprof`
   * Benchmarking with `testing.B`

3. **Configuration & env**

   * Use `flag` and `os` packages
   * Reload configs without downtime

4. **Error handling**

   * Idiomatic error wrapping (`errors`, `%w`)
   * Contextual error messages in network code

---

## Capstone Projects

Once you’re comfortable, build these mini-projects:

1. **Custom reverse proxy**

   * Parse config file with backends
   * Load-balance across multiple servers
   * Add request logging + metrics

2. **HTTP/2 file server**

   * Serve static files efficiently
   * Experiment with multiplexing & TLS

3. **Tiny service mesh component**

   * Route requests based on host/path
   * Add retry & circuit breaker middleware
   * Inspired by Traefik/Envoy
