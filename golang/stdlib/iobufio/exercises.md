# Exercises - IO package

### **A. Core Concepts (io.Reader / io.Writer Basics)**

1. Implement a program that reads a string from `strings.NewReader` and prints it using `io.Copy`.
2. Create a custom type that implements `io.Reader` and always returns `"Hello"` regardless of input.
3. Create a custom type that implements `io.Writer` and writes to `os.Stdout` with a prefix `[LOG]`.
4. Use `io.CopyN` to copy only the first 10 bytes from a reader to a writer.
5. Use `io.LimitReader` to restrict reading from a file to a specific byte count.
6. Chain multiple readers using `io.MultiReader` and print combined output.
7. Chain multiple writers using `io.MultiWriter` to write to both a file and `os.Stdout` simultaneously.
8. Measure how many bytes were copied using `io.Copy` between two readers/writers.
9. Wrap a reader with `io.TeeReader` to print content while it’s being read and copied elsewhere.
10. Compare two streams for equality using `io.ReadAll` and byte comparison.

### **B. File I/O Practice**

11. Read an entire text file using `io.ReadAll` and print the contents.
12. Write data to a file using a custom `io.Writer` that appends timestamps to each line.
13. Implement a file copier using `io.Copy` that copies one file to another.
14. Use `io.LimitReader` to copy only a portion of a file into another.
15. Combine multiple small files into one large file using `io.MultiReader`.
16. Write a function that logs file read/write byte count using `io.Copy` and `io.Pipe`.
17. Build a command-line tool that mimics `cat` using `io.Copy(os.Stdout, file)`.
18. Implement a function that reverses the content of a file (line by line) using `io.Reader` and buffers.
19. Create a reader that skips the first N bytes before returning content (custom wrapper around another reader).
20. Build a small program that reads binary data from a file and interprets it as structured bytes (for network packet simulation).

### **C. Networking-Oriented Exercises**

21. Implement a TCP echo server that uses `io.Copy(conn, conn)` for bidirectional data flow.
22. Create a TCP client that connects to the echo server and sends input from `os.Stdin` using `io.Copy`.
23. Use `io.Copy` to forward data between two TCP connections (simple proxy).
24. Build a simple HTTP file server that serves files using `io.Copy` from file to response writer.
25. Read request body in an HTTP handler using `io.ReadAll` and log the size.
26. Implement a bandwidth logger that wraps a TCP connection reader/writer and counts total bytes transferred.
27. Build a TCP relay: read from one connection, write to another using `io.Copy` and `io.Pipe`.
28. Use `io.TeeReader` in a TCP server to log incoming data to a file while forwarding it to another destination.
29. Build a function that transfers data from a UDP connection to stdout using `io.Copy`.
30. Implement a simple chat relay server using `io.Copy` between multiple TCP clients.

### **D. Piping and Streaming Exercises**

31. Use `io.Pipe` to connect a writer goroutine and a reader goroutine — simulate a stream pipeline.
32. Chain transformations: Reader → Uppercase Filter → Writer using custom Reader wrappers.
33. Use `io.Pipe` to simulate a file upload process: writer writes chunks, reader simulates receiving.
34. Create a gzip compressor using `io.Pipe` + `compress/gzip` + `io.Copy`.
35. Create a line counter that wraps any `io.Reader` and counts newlines while passing data through.
36. Combine `io.TeeReader` and `io.Pipe` to split incoming stream into processing and logging paths.
37. Build a JSON streaming parser that uses `io.Reader` and `json.Decoder` to handle long data streams.
38. Use `io.CopyBuffer` with a manual buffer to control memory usage during file transfer.
39. Implement a simple “stream duplicator” that reads from stdin and writes to two network connections using `io.MultiWriter`.
40. Build a command that reads logs from stdin, filters them line by line, and writes filtered output to both file and network using `io.MultiWriter`.

---

# Exercises - Bufio package

### **A. Buffered Reader Fundamentals**

1. Read lines from a text file using `bufio.NewReader` and `ReadString('\n')`, print each line.
2. Use `bufio.Reader.ReadBytes('\n')` to process binary data containing newline-terminated records.
3. Compare performance of reading a large file with and without buffering (`io.ReadAll` vs `bufio.Reader`).
4. Use `Peek(n)` to preview upcoming bytes in a file without consuming them.
5. Implement a simple tokenizer using `ReadBytes(' ')` to split input into words.
6. Write a function that reads until a specific delimiter (e.g., `"END"`) using `ReadSlice`.
7. Build a custom `ReadLine` loop using `bufio.Reader.ReadLine()` that handles long lines.
8. Combine `bufio.Reader` with a network `net.Conn` to efficiently read client messages in a TCP server.
9. Create a buffered stdin reader and read user commands interactively (`ReadString('\n')`).
10. Wrap a `bytes.Buffer` with `bufio.NewReader` and measure how `Buffered()` changes after reads.

### **B. Buffered Writer Fundamentals**

11. Write a program that logs messages to a file using `bufio.NewWriter`, flushing only every 5 writes.
12. Implement a buffered logger that automatically flushes when buffer reaches a threshold (`Available()` < N).
13. Use `bufio.Writer.WriteString` to write multiple strings efficiently to a network socket.
14. Compare latency of writing small messages with and without `bufio.Writer` to a TCP connection.
15. Build a function that writes JSON lines to a file using `bufio.NewWriter` and flushes every N objects.
16. Use `Flush()` in a deferred call to ensure data integrity even if an error occurs mid-write.
17. Implement a file copier that reads with `bufio.Reader` and writes with `bufio.Writer`, testing throughput.

### **C. Scanner and Advanced Usage**

18. Use `bufio.Scanner` to read words from a file with `ScanWords`, count total words.
19. Write a program that scans CSV lines and prints fields using `bufio.Scanner` with a custom split function.
20. Build a network line-based protocol parser (like Redis protocol) using `bufio.Reader` and `ReadString('\r')`.

**Coverage Summary:**

!!! note
    These exercises bridge `io` fundamentals with real buffering control — the step between raw I/O and efficient, production-grade stream handling.
