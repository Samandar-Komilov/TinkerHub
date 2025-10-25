# Roadmap: `net/http` (trivial → advanced)

### 0. Absolute basics — `Request`, `Response`, `Server`, `ResponseWriter`, `ListenAndServe`, `Serve`

**APIs:** `http.Handler` / `http.HandlerFunc`, `http.ListenAndServe`, `http.Serve`, `http.ResponseWriter`, `*http.Request`. ([Go Packages][1])
**Microproject:** Minimal server with `/health` and `/item/{id}` endpoints. Return JSON and proper status codes.

**Minimal server**

```go
http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","text/plain")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("ok"))
})
log.Fatal(http.ListenAndServe(":8080", nil))
```

### 1. Request Context - `context.Context`

**APIs:** 
**Microproject:** 


### 2. Routing: Standard Library (`ServeMux`) and Third-Party (`chi`, `gorilla/mux`)

**APIs to inspect:** `http.ServeMux`, path parameters as `r.Context()` values (chi stores params in context).
**Microproject:** Implement a tiny mux: map path → `http.Handler` and support path params (e.g., `/item/:id`), then compare to `chi` usage.

### 3. Middleware & handler composition (practical patterns)

**APIs:** `http.HandlerFunc`, wrapper functions `func(http.Handler) http.Handler`, `context` for request-scoped data.
**Microproject:** Build a chainable middleware set: logging, request ID (X-Request-ID), panic recovery, and timeout middleware that attaches a context with deadline to `r`.
**What you learn / gotchas:** Middleware order matters. Prefer `Request.WithContext(ctx)` for cancellation and pass data through `context` (but don’t use context for business data). Keep middleware cheap and non-blocking. Example: a logging middleware must not `Read` the request body (it will interfere with handler).

**Example logging middleware skeleton**

```go
func Logging(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    // wrap ResponseWriter if you need status/bytes
    next.ServeHTTP(w, r)
    log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
  })
}
```

### 4. Request body, streaming and large uploads

**APIs:** `r.Body (io.ReadCloser)`, `io.Copy`, `http.MaxBytesReader`, `multipart.Reader`.
**Microproject:** Implement a streaming upload endpoint that writes to a temp file while enforcing a max size and protecting against slow clients.
**What you learn / gotchas:** Don’t `ReadAll` large bodies. Use `MaxBytesReader` to protect memory. For uploads stream straight to disk. Server auto-closes `r.Body` after the handler returns — don’t `defer r.Body.Close()` in the handler (server covers it). If you hijack the connection, you become responsible for closing/cleanup.

---

### 5. HTTP client basics — `http.Client`, `Do`, `NewRequest`

**APIs:** `http.Client`, `http.NewRequest`, `Client.Do`, `http.Get/Post`.
**Microproject:** Implement a single-feed fetcher that uses `NewRequestWithContext(ctx, "GET", url, nil)` and sets `Accept` headers for RSS/XML.
**What you learn / gotchas:** Use `Request.WithContext` to propagate cancellation. **Always** `defer resp.Body.Close()` on client responses. Small mistake: creating a new `http.Client` (or `Transport`) per request kills connection reuse and throughput — reuse clients/transport between requests. ([Stack Overflow][2])

**Example fetch**

```go
ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
defer cancel()
req, _ := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
req.Header.Set("Accept","application/rss+xml, application/xml; q=0.9, */*;q=0.1")
resp, err := client.Do(req)
if err != nil { return err }
defer resp.Body.Close()
body, _ := io.ReadAll(resp.Body)
```

---

### 6. Transport & connection reuse — real performance knobs

**APIs:** `http.Transport` (fields: `MaxIdleConns`, `MaxIdleConnsPerHost`, `IdleConnTimeout`, `TLSHandshakeTimeout`, `DisableCompression`, `MaxConnsPerHost`, `ForceAttemptHTTP2`).
**Microproject:** Build a reusable `http.Client` with a tuned `Transport` and use it to fetch 100 feeds concurrently (bounded concurrency — see step 9). Measure throughput.
**What you learn / gotchas:** Transport caches connections; reusing it is essential for speed. Tuning `MaxIdleConnsPerHost` and `IdleConnTimeout` matters for many hosts. By default the Transport may request gzip and transparently decompress responses (you can change this via `DisableCompression`). Do **not** create a fresh Transport per request. ([Go][3])

**Good client template**

```go
tr := &http.Transport{
  MaxIdleConns:        100,
  MaxIdleConnsPerHost: 20,
  IdleConnTimeout:     90 * time.Second,
  TLSHandshakeTimeout: 10 * time.Second,
  // ForceAttemptHTTP2: true, // optional
}
client := &http.Client{Transport: tr, Timeout: 15 * time.Second}
```

---

### 7) Timeouts, cancellation and deadlines (defensive programming)

**APIs:** `req.WithContext(ctx)`, `http.Client.Timeout`, `Dialer.Timeout`, `Transport.TLSHandshakeTimeout`.
**Microproject:** Implement a fetcher that enforces a per-feed timeout (e.g., 10s) and an overall aggregator timeout (e.g., 60s). Cancel outstanding fetches when overall timeout hits.
**What you learn / gotchas:** Prefer per-request contexts to blanket `Client.Timeout` if you need nuanced cancellation (read vs connect vs total). If you use both, be careful how they interact; `Client.Timeout` cancels the whole operation. Always design for partial results and graceful cancellation.

---

### 8) Compression, encodings and conditional requests

**APIs:** `Request.Header` (If-Modified-Since, If-None-Match), `resp.Header.Get("ETag")`, `Transport.DisableCompression`.
**Microproject:** Implement conditional fetching for feeds: store `ETag` and `Last-Modified` for each feed and send conditional GETs to avoid downloading unchanged feeds (handle `304 Not Modified`). Also accept gzipped content and handle it correctly.
**What you learn / gotchas:** Transport can automatically add `Accept-Encoding: gzip` and transparently decompress the body for you unless you disable that behavior. Using conditional GETs saves bandwidth and reduces parsing cost. ([go.googlesource.com][4])

---

### 9) Concurrency control, retries, backoff, and politeness

**APIs / primitives:** goroutines + bounded worker pool (channel semaphore), retry strategy (exponential backoff), `RateLimiter` (token bucket).
**Microproject:** Build a bounded concurrent fetcher: N workers pull URLs from a queue, each fetch uses the shared `client`, applies retries with exponential backoff and respects per-host rate limits. Add per-host semaphore if you must limit concurrency per origin.
**What you learn / gotchas:** Unbounded concurrency destroys CPU, memory and remote servers. Respect remote servers: use `Retry-After` headers and backoff. Keep retries idempotent (GET is safe).

---

### 10) Tracing & debugging network problems with `httptrace`

**APIs:** `net/http/httptrace.ClientTrace`, `httptrace.WithClientTrace`.
**Microproject:** Add an instrumentation mode to your fetcher that prints timing for DNS, TCP connect, TLS handshake and first byte for each feed — use `httptrace` to collect these. This is how you find where slowdowns (DNS vs connect vs server) happen. ([Go][5])

**Short httptrace snippet**

```go
trace := &httptrace.ClientTrace{
  DNSStart: func(_ httptrace.DNSStartInfo){ log.Println("dns start") },
  ConnectStart: func(network, addr string){ log.Println("connect", network, addr) },
  GotFirstResponseByte: func(){ log.Println("first byte") },
}
req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
resp, err := client.Do(req)
```

---

### 11) Testing: `httptest`, `ResponseRecorder`, `httptest.NewServer`

**APIs:** `net/http/httptest` (ResponseRecorder, NewRequest, NewServer).
**Microproject:** Unit-test your aggregator handlers using `httptest.NewRecorder()` for handler logic and `httptest.NewServer()` for end-to-end client tests (mock remote feed responses and simulate 200/304/gzip/slow responses).
**What you learn / gotchas:** `ResponseRecorder` lets you assert on status/body without a real network. `httptest.NewServer` gives a real HTTP URL useful for exercising the real `http.Client` and Transport behavior. Use both appropriately. ([Go Packages][6])

---

### 12) Advanced: reverse proxy, `httputil.ReverseProxy`, connection lifecycle hooks, graceful shutdown, metrics

**APIs:** `httputil.ReverseProxy`, `Server.ConnState`, `Server.Shutdown`, `Server.BaseContext` / `ConnContext`.
**Microproject:** Implement a reverse proxy endpoint that forwards requests to a feed-fetch service and attaches a short cache; implement graceful shutdown so in-flight fetches can finish and new accepts stop. Add a `/metrics` endpoint for Prometheus.
**What you learn / gotchas:** `Shutdown(ctx)` triggers graceful stop but you must manage long-running handlers. `ConnState` and `ConnContext` are advanced tools when you need per-connection bookkeeping. HTTP/2 is enabled automatically by `net/http` in modern Go — but you can adjust it via `TLSNextProto` / `Transport.ForceAttemptHTTP2` if you need special behavior. ([Go Packages][7])

---

# RSS-Aggregator specific checklist (immediately actionable)

1. **One shared `http.Client`** with tuned `Transport` (reused globally). Never create per-request clients. ([Stack Overflow][2])
2. **Bound concurrency**: limit simultaneous fetches (e.g., 50 global; per-host 2).
3. **Per-request timeout** via `req.WithContext(context.WithTimeout(...))`.
4. **Conditional GETs** using `ETag` / `If-Modified-Since`. Store ETag/Last-Modified per feed.
5. **Respect compression**: let Transport decompress, or handle it if you need the raw bytes. ([go.googlesource.com][4])
6. **Instrumentation**: add `httptrace` hooks to see slow DNS / slow TLS / slow server. ([Go][5])
7. **Testing**: mock slow/erroneous feeds with `httptest.NewServer()` and test cancellation and retry logic. ([Go Packages][6])

---

# Common mistakes (be blunt)

* Creating `http.Client`/`Transport` per request — kills connection reuse and will exhaust sockets. Fix: one shared client/transport. ([Stack Overflow][2])
* Forgetting `resp.Body.Close()` — leaks FDs and prevents connection reuse. Always `defer resp.Body.Close()` or `io.Copy(io.Discard, resp.Body)` then close when you drop content.
* Unbounded goroutines for fetches — memory and CPU explosions. Use a bounded worker pool.
* Ignoring conditional GETs — you waste bandwidth and CPU parsing identical feeds.
* Not tracing — when feed fetching is slow, you won’t know if it’s DNS, TCP, TLS, or server slowness without `httptrace`.

---

# Testing & debugging recipes (practical)

* To test handler logic: `httptest.NewRecorder()` + `httptest.NewRequest()`. ([Go Packages][6])
* To debug a slow feed fetch: attach `httptrace.ClientTrace` and print DNS/Connect/TLS/FirstByte timings. ([Go][5])
* To validate transport reuse: run N requests and inspect `GODEBUG=http2debug=1` or use `httptrace` to see new connections.
