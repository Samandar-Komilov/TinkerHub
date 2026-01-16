# Zig Web Servers: A Concurrency Exploration

This project contains three simple web servers written in Zig, each demonstrating a different model for handling concurrent connections.

- `mproc_srv.zig`: Multi-process server (fork-per-connection)
- `mthread_srv.zig`: Multi-threaded server (thread-per-connection)
- `async_srv.zig`: Asynchronous server (event-driven with async/await)

To build and run any of these servers:

```sh
zig build-exe <filename>.zig
./<filename>
```

For example:
```sh
zig build-exe mproc_srv.zig
./mproc_srv
```

You can then connect to the server using `curl` or a web browser:
- `mproc_srv`: `curl http://localhost:8080`
- `mthread_srv`: `curl http://localhost:8081`
- `async_srv`: `curl http://localhost:8082`

## Core Networking Concepts

All three servers use the same fundamental networking steps, provided by Zig's standard library (`std.net`):

1.  **Socket Creation**: A socket is the endpoint for network communication. `std.net.StreamServer.listen()` creates a TCP socket for us.
2.  **Binding**: The socket is associated with an IP address and a port number. `listen()` handles this, binding to a port on all available network interfaces.
3.  **Listening**: The server starts listening for incoming connections on the bound socket. `listen()` also marks the socket as a passive socket that can accept incoming connection requests.
4.  **Accepting**: When a client connects, the server accepts the connection. This creates a *new* socket dedicated to communication with that specific client. The original listening socket remains open to accept more connections. In our code, `stream_server.accept()` performs this step, blocking until a connection arrives.

## Concurrency Models

The interesting part is how each server handles multiple clients at the same time.

### 1. Multi-Process Server (`mproc_srv.zig`)

This is a classic UNIX model.

-   **Concept**: When a new connection is accepted, the main server process creates an identical copy of itself using `fork()`.
-   **Zig Feature**: `std.os.fork()`
-   **How it works**:
    -   The parent process's only job is to accept new connections and fork. It immediately closes its copy of the connection socket and goes back to waiting for the next connection.
    -   The child process inherits the connection socket. It handles the request (sends the HTTP response), closes its connection, and then exits.
-   **Pros**: Excellent isolation. A crash in a child process will not affect the main server or other connections.
-   **Cons**: `fork()` can be relatively slow and memory-intensive, as it involves copying the process's address space (though modern OSes use copy-on-write to optimize this). It's less common for high-performance servers today.

### 2. Multi-Threaded Server (`mthread_srv.zig`)

A very common and straightforward concurrency model.

-   **Concept**: Instead of a new process, the server spawns a new thread for each connection.
-   **Zig Feature**: `std.Thread.spawn()`
-   **How it works**:
    -   The main thread accepts a connection and immediately hands it off to a new thread.
    -   The new thread is responsible for handling the request and closing the connection.
    -   Threads share the same memory space, which makes them more lightweight than processes.
-   **Pros**: Faster to create threads than processes. Shared memory can make communication between threads easier (though not used in this simple example).
-   **Cons**: Less isolation than processes. An unhandled error (like a crash) in one thread can take down the entire server process. There can be overhead from context-switching between many threads.

### 3. Asynchronous Server (`async_srv.zig`)

This model uses an event loop and non-blocking I/O to handle many connections within a single thread (or a small number of threads).

-   **Concept**: Instead of blocking on I/O operations (like reading a request), the server registers its interest in an event (e.g., "data available to read") and then moves on to do other work. An event loop notifies the server when the event occurs.
-   **Zig Features**: `async`, `await` (though `await` is not explicitly used in this simplified example).
-   **How it works**:
    -   The main loop accepts a connection.
    -   It calls `async handleConnection(...)`. This doesn't run the function immediately. Instead, it creates an "async frame" (a coroutine or "green thread") and schedules it to be run.
    -   The Zig runtime's scheduler manages a pool of threads and runs these async frames on them.
    -   This provides concurrency without the overhead of creating a full OS thread for every single connection. The scheduler can run many async frames on one OS thread.
-   **Note on "True" Async**: In this specific example, the `reader.read()` call inside `handleConnection` is still *blocking*. A fully optimized async server would use `await` on non-blocking I/O operations. This would allow the OS thread to run other async frames while waiting for data, leading to much better performance under high load. However, this example still effectively demonstrates how `async` is used to achieve concurrency.
-   **Pros**: Extremely efficient and scalable. Can handle tens of thousands of connections with very few OS threads. Low memory usage per connection.
-   **Cons**: Can lead to more complex "callback-style" or `async/await`-heavy code. Debugging can be trickier than in a simple blocking model.
