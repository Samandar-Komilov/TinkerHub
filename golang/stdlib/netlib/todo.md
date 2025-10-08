# Step-by-step learning path (trivial → advanced)

### 1) Minimal TCP client & server (the absolute baseline)

* **APIs:** `net.Listen`, `net.Listener.Accept`, `net.Dial`, `net.Conn.Read/Write/Close`.
* **Microproject:** TCP echo server + client that exchanges newline-terminated strings. Use `bufio.Reader.ReadString('\n')` on the server.
* **What you learn / gotchas:** Accept loop → per-connection goroutine. Always `defer conn.Close()` inside handler. Use `bufio` to avoid tiny reads. Understand that `net.Conn` gives you stream semantics (like `read()`/`write()` in C). See `net.Conn` methods (`SetDeadline`, etc.) below. ([Go][2])
* **Blog heading:** “TCP from scratch: an echo server and what `net.Conn` actually exposes”

**Minimal skeleton**

```go
// server.go
ln, _ := net.Listen("tcp", "127.0.0.1:6000")
defer ln.Close()
for {
  conn, _ := ln.Accept()
  go func(c net.Conn){
    defer c.Close()
    r := bufio.NewReader(c)
    for {
      s, err := r.ReadString('\n')
      if err != nil { return }
      c.Write([]byte("echo: "+s))
    }
  }(conn)
}
```

---

### 2) Concurrency patterns & graceful shutdown

* **APIs:** `Listener.Close()`, accept loop patterns, `sync.WaitGroup`, cancellation with `context`.
* **Microproject:** Make the echo server support graceful shutdown (stop accepting, close active conns, wait for handlers).
* **Gotchas:** `Accept()` unblocks with `Listener.Close()`; closing a `net.Conn` will unblock `Read()` in another goroutine. Don’t leak goroutines.

---

### 3) Deadlines & timeouts (practical defense)

* **APIs:** `conn.SetDeadline`, `conn.SetReadDeadline`, `conn.SetWriteDeadline`.
* **Microproject:** Make the server enforce a 30s idle timeout per connection; client demonstrates timed read/write failure.
* **Gotchas:** Deadlines are per-call and racey if you set them while a call is already blocked; prefer setting before blocking calls. See `net.Conn` docs for exact semantics. ([Go][2])

---

### 4) TCP specifics: `TCPConn`, `ListenTCP`, connection shutdown

* **APIs:** `net.TCPConn` methods: `CloseRead`, `CloseWrite`, `SetKeepAlive`, `SetLinger`, `File()` (careful).
* **Microproject:** Build a simple TCP proxy that forwards bytes between client and backend (full duplex).
* **Gotchas:** `CloseRead`/`CloseWrite` map directly to `shutdown()` halves. `File()` gives a duped FD — closing behavior differs. You’ll hit keepalive platform differences; Go’s defaults are not always what you expect. ([Go][3])

---

### 5) Dialer, DialContext, and fine control of outgoing connections

* **APIs:** `net.Dialer`, `Dialer.DialContext`, `Dialer.Timeout`, `Dialer.Control`.
* **Microproject:** Client that uses `Dialer.DialContext` with a 3s timeout and a `Control` callback to set a socket option (example: `SO_REUSEADDR` or binding source address).
* **Why it matters:** `Dialer` gives cancellation (via `context`) and a hook to configure socket options before connect. This is how you implement custom source addresses, local ports, and some per-socket options. ([Go][4])

**Dialer sketch**

```go
d := net.Dialer{Timeout: 3*time.Second, Control: func(network, address string, rc syscall.RawConn) error {
  var controlErr error
  rc.Control(func(fd uintptr){ controlErr = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1) })
  return controlErr
}}
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()
conn, err := d.DialContext(ctx, "tcp", "1.2.3.4:80")
```

---

### 6) UDP & `PacketConn` (connectionless)

* **APIs:** `net.ListenPacket`, `net.PacketConn.ReadFrom` / `WriteTo`, `net.DialUDP`, `net.UDPConn`.
* **Microproject:** UDP echo server that echos back the sender address and implements a small request ID to handle message reordering.
* **What’s different from TCP:** Each `ReadFrom` returns a discrete packet with a source address; no stream framing. Buffer sizing matters — you need to provision for worst-case packet size. ([Go][5])

**UDP skeleton**

```go
pc, _ := net.ListenPacket("udp", ":6001")
defer pc.Close()
buf := make([]byte, 2048)
n, addr, _ := pc.ReadFrom(buf)
pc.WriteTo([]byte("echo:"+string(buf[:n])), addr)
```

---

### 7) Raw IP / ICMP & advanced packet control

* **APIs:** `net.ListenIP`, `net.ListenPacket` for raw/IP, plus `golang.org/x/net/ipv4` and `ipv6` for control messages.
* **Microproject:** Build a tiny `ping` (ICMP echo) client and explain why on many systems you need elevated privileges or a special per-OS path.
* **Gotchas:** Raw sockets require privileges and are system-dependent. Use `x/net/ipv4` for ancillary control messages (TTL, TOS, IP options). See `ListenPacket` docs for ICMP caveats. ([Go Packages][6])

---

### 8) Unix domain sockets (IPC)

* **APIs:** `net.Listen("unix", path)`, `net.Dial("unix", path)`, `net.UnixConn`.
* **Microproject:** HTTP server that listens on a unix socket and an HTTP client that connects to it (show `http.Serve` with a custom `net.Listener`).
* **Why use it:** Lower latency on same-host comms, file permissions based access control, and commonly used in containerized apps for sidecars.

---

### 9) Name resolution & `net.Resolver`

* **APIs:** `net.LookupIP`, `net.LookupHost`, `net.Resolver` (with custom `Dial` and `PreferGo`).
* **Microproject:** Implement a small client that uses a custom `net.Resolver` with a `Dialer` to query a chosen DNS server (useful for testing and DoH/DoT wrappers).
* **Gotchas:** Go has two resolution modes (Go resolver vs cgo). Behavior differs by OS and can affect blocking/cancellations — prefer `Resolver` with contexts if you need timeouts/ cancelation. (There are platform nuances; test on your target OS.) ([Go][7])

---

### 10) Integration: using `net` with `crypto/tls` and `net/http`

* **APIs:** `crypto/tls.Client/Server` accept `net.Conn`, `http.Server.Serve(listener)`, custom `Transport.DialContext`.
* **Microproject:** Handcraft a TLS client using `tls.Client` over a raw `net.Conn`, and show a custom `http.Transport` that sets a `DialContext` for connection reuse/troubleshooting.
* **Why:** TLS and `net/http` are thin layers on `net` — learning `net` means you can control connection lifecycles, timeouts, and tracing at the socket level. (Go’s `crypto/tls` and `net/http` are designed to interoperate with `net.Conn`.) ([Go Packages][1])

---

### 11) Syscall hooks & socket options (advanced)

* **APIs:** `Conn.(syscall.Conn).SyscallConn()`, `Dialer.Control`, `syscall.RawConn.Control`.
* **Microproject:** Implement setting `SO_REUSEPORT`/`SO_BINDTODEVICE` on outgoing/incoming sockets (platform-specific).
* **Gotchas:** This is platform and kernel dependent. You will need to use `syscall` and respect that some options require binding before `connect()`. Use `Dialer.Control` to set options between socket creation and connect. ([Stack Overflow][8])

---

### 12) Testing patterns: `net.Pipe`, in-memory testing

* **APIs:** `net.Pipe()`.
* **Microproject:** Unit test a protocol handler using `net.Pipe` to create two in-memory connected `net.Conn`s. Use this to show how to test timeouts and backpressure deterministically.
* **Why:** `net.Pipe` is perfect for unit tests that must avoid real network nondeterminism. ([Go][9])

---

# Compact cheat-sheet (what to remember)

* `net.Dial` / `net.Dialer` / `DialContext` — client side (timeouts, control hooks). ([Go][4])
* `net.Listen` / `Listener.Accept` — server side.
* `net.Conn` — `Read/Write/Close/LocalAddr/RemoteAddr/SetDeadline`. ([Go][2])
* `net.TCPConn`, `net.UDPConn`, `net.UnixConn` — protocol-specific extra methods.
* `net.PacketConn` (`ListenPacket`) — UDP and datagram semantics. ([Go][5])
* `net.Resolver` — controlled DNS with `context` and custom Dial.
* `syscall.RawConn` via `SyscallConn()` and `Dialer.Control` — set socket options. ([Stack Overflow][8])
* `net.Pipe()` for in-memory connection testing. ([Go][9])

---

# Pitfalls you will hit (be explicit)

1. **DNS blocking vs cancellation** — default resolver paths differ across OSes; use `Resolver` + `Context` if you need timeouts. ([Go][7])
2. **Deadlines race** — setting deadlines while a read is in progress has subtle semantics; prefer setting deadlines before entering blocking ops. ([Go][2])
3. **Socket option timing** — some options must be set before `connect()`; use `Dialer.Control`. ([Stack Overflow][8])
4. **UDP is packet-oriented** — you must handle packet boundaries yourself; a single `ReadFrom` reads one datagram. ([Go][5])
5. **Permissions for raw sockets** — ICMP/raw sockets need privileges; testing may require `CAP_NET_RAW` or similar. ([Go Packages][6])

---

# Blog post skeleton (how to turn this into one post)

* Title: “Practical tour of Go’s `net` package — from `Dial` to raw sockets”
* TL;DR: one-paragraph: what you’ll build and the promise (learn by coding).
* Intro: why `net` is worth learning even if you know sockets in C (it’s not magic — it’s the same pieces, fewer footguns). Cite `net` overview. ([Go Packages][1])
* Sections: follow the step order above. Each section: *Problem statement → Minimal code → What to watch for (gotchas) → When to use in production*.
* Appendices: cheat-sheet, references to relevant `x/net` packages (`ipv4/ipv6/icmp`), and small benchmark notes.
* Closing: “When to stop reimplementing” — recommend when to stop and use `net/http`, `crypto/tls`, or a third-party lib.

---

# Further reading & authoritative refs

* Official `net` package overview and API. ([Go Packages][1])
* `net.Conn` source / docs (methods and deadlines). ([Go][2])
* `net.Dialer` / `DialContext` implementation & notes. ([Go][4])
* UDP/`PacketConn` (UDP/ListenPacket/ReadFrom/WriteTo). ([Go][5])
* `net.Pipe()` implementation and testing usage. ([Go][9])

