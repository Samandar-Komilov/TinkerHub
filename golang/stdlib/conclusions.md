# Go Standard Library - Conclusions

### io

The `io` package provides universal abstractions for input/output streams - files, network and database sockets, pipes, in-memory buffers, strings, etc. We don't care if its socket or file, a single interface provides functionality to work with them.

1. `io.Reader` - abstraction for all data sources. Anything that can produce bytes implement this interface:
    ```go
    type Reader interface {
        Read(p []byte) (n int, err error) // p - destination buffer, n - number of bytes read
    }
    ```
    Here is a simple example of reading from string stream:
    ```go
    func main() {
        r := strings.NewReader("Hello World!")
        buf := make([]byte, 4)

        for {
            n, err := r.Read(buf)
            if n > 0 {
                fmt.Printf("%d", buf[:n])
            }

            if err == io.EOF {
                break
            }
        }

        fmt.Println("We have just read from string stream!")
    }
    ```

2. `io.Writer` - abstraction for all data receivers. Anything that can receive bytes implements this interface.
    ```go
    type Writer interface {
        Write(p []byte) (n int, err error) // p - data being sent in bytes, n - number of bytes written
    }
    ```
    Here is a simple example of writing to a file:
    ```go
    func WriterExample() {
        f, _ := os.Create("hello.txt")
        defer f.Close()
        f.Write([]byte("Hello World!"))
        fmt.Println("File write completed!")
    }
    ```

3. `io.Copy(dst, src)` - universal data pump. This function copies from any `Reader` to any `Writer` until `EOF` or error occurs.
    ```go
    func Copy(dst Writer, src Reader) (written int64, err error) {
        return copyBuffer(dst, src, nil)
    }
    ```
    Here is a simple example of file copy:
    ```go
    func CopyExample() {
        src, _ := os.Open("files/in.txt")
        dst, _ := os.Create("files/out.txt")
        defer src.Close()
        defer dst.Close()
        io.Copy(dst, src)
    }
    ```

4. `io.CopyN(dst, src, n)` - limited copy. Used to copy a fixed number of bytes from source to destination.
    ```go
    func CopyN(dst Writer, src Reader, n int64) (written int64, err error) {
        written, err = Copy(dst, LimitReader(src, n))
        if written == n {
            return n, nil
        }
        if written < n && err == nil {
            // src stopped early; must have been EOF.
            err = EOF
        }
        return
    }
    ```
    The example is the same as (4), just we copy N bytes from source file to destination file.

5. `io.CopyBuffer(src, dst, buf)` - use your own buffer for performance tuning (e.g. memory reuse).
    ```go
    func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
        if buf != nil && len(buf) == 0 {
            panic("empty buffer in CopyBuffer")
        }
        return copyBuffer(dst, src, buf)
    }
    ```
    The `copyBuffer()` internal function copies the source to the destination using our buffer inside infinite loop until it reaches `EOF` or error.

    The same logic, but this time we give our buffer:
    ```go
    // ...
    buffer := make([]byte, 5)
	io.CopyBuffer(dst, src, buffer)
    ```

6. `io.LimitReader(r, n)` - wraps a Reader but stops after n bytes. Prevents over-reading large inputs (e.g. user uploads).
    ```go
    func LimitReader(r Reader, n int64) Reader { return &LimitedReader{r, n} }

    type LimitedReader struct {
        R Reader // underlying reader
        N int64  // max bytes remaining
    }
    ```
    Here is a simple example:
    ```go
    r := strings.NewReader("ABCDEFGHIJK")
    limited := io.LimitReader(r, 5)
    io.Copy(os.Stdout, limited) // prints ABCDE
    ```

7. `io.MultiReader(r1, r2, ...)` - reading from chained data sources. Sequentially reads from multiple readers.
    ```go
    type multiReader struct {
        readers []Reader
    }

    func (mr *multiReader) Read(p []byte) (n int, err error) {
        for len(mr.readers) > 0 {
            // Optimization to flatten nested multiReaders (Issue 13558).
            if len(mr.readers) == 1 {
                if r, ok := mr.readers[0].(*multiReader); ok {
                    mr.readers = r.readers
                    continue
                }
            }
            n, err = mr.readers[0].Read(p)
            if err == EOF {
                // Use eofReader instead of nil to avoid nil panic
                // after performing flatten (Issue 18232).
                mr.readers[0] = eofReader{} // permit earlier GC
                mr.readers = mr.readers[1:]
            }
            if n > 0 || err != EOF {
                if err == EOF && len(mr.readers) > 0 {
                    // Don't return EOF yet. More readers remain.
                    err = nil
                }
                return
            }
        }
        return 0, EOF
    }

    func MultiReader(readers ...Reader) Reader {
        r := make([]Reader, len(readers))
        copy(r, readers)
        return &multiReader{r}
    }
    ```
    As we can see, the official code loops over the slice of readers and copies from each sequentially.  
    Here is a simple example:
    ```go
    func MultiReaderExample() {
        r := io.MultiReader(strings.NewReader("Hello"), strings.NewReader("World!"))
        io.Copy(os.Stdout, r)
    }
    ```

8. `io.MultiWriter(w1, w2, ...)` - broadcast data to multiple sources. For example, write to a file, stdout and network simultaneously.
    ```go
    type multiWriter struct {
        writers []Writer
    }

    func (t *multiWriter) Write(p []byte) (n int, err error) {
        for _, w := range t.writers {
            n, err = w.Write(p)
            if err != nil {
                return
            }
            if n != len(p) {
                err = ErrShortWrite
                return
            }
        }
        return len(p), nil
    }

    func MultiWriter(writers ...Writer) Writer {
        allWriters := make([]Writer, 0, len(writers))
        for _, w := range writers {
            if mw, ok := w.(*multiWriter); ok {
                allWriters = append(allWriters, mw.writers...)
            } else {
                allWriters = append(allWriters, w)
            }
        }
        return &multiWriter{allWriters}
    }
    ```
    Here is a simple example:
    ```go
    f, _ := os.Create("log.txt")
    mw := io.MultiWriter(os.Stdout, f)
    mw.Write([]byte("Hello")) // prints and logs
    ```

9. `io.TeeReader(r, w)` - side channel logging. Reads from `r`, writes every byte to `w` as it passes through. Used to inspect or log the data while streaming to somewhere.
    ```go
        func TeeReader(r Reader, w Writer) Reader {
        return &teeReader{r, w}
    }

    type teeReader struct {
        r Reader
        w Writer
    }

    func (t *teeReader) Read(p []byte) (n int, err error) {
        n, err = t.r.Read(p)
        if n > 0 {
            if n, err := t.w.Write(p[:n]); err != nil {
                return n, err
            }
        }
        return
    }
    ```
    Here is a simple example:
    ```go
    func TeeReaderExample() {
        src := strings.NewReader("secret data")
        log := &bytes.Buffer{}
        tee := io.TeeReader(src, log)
        io.Copy(io.Discard, tee)
        fmt.Println("Logged:", log.String(), log)
    }
    ```

10. `io.ReadAll(r)` - conveniently read full data. Useful for small, complete reads (configs, request bodies).  
    Here is an example:
    ```go
    func ReadAllExample() {
        r := strings.NewReader("all at once")
        data, _ := io.ReadAll(r)
        fmt.Println(string(data))
    }
    ```

11. `io.Pipe()` - in-memory stream connection. Connects a writer and reader through a pipe (synchronized). It acts as a bridge between goroutines - one produces and one consumes.  
    Here is a simple example:
    ```go
    r, w := io.Pipe()
    go func() {
        fmt.Fprint(w, "streamed data")
        w.Close()
    }()
    io.Copy(os.Stdout, r)
    ```
    This is not the OS level pipe that is used to intercommunicate over processes. This is a kind of channel which is used to communicate between goroutines.

12. `io.NopCloser(r)` - add `Close` to Reader. It is necessary when a function or interface expects an `io.ReadCloser` (which includes a `Close()` method), but the underlying data source is an `io.Reader` that does not require or benefit from a `Close()` operation.  
    Here is an example:
    ```go
    func processData(r io.ReadCloser) error {
        defer r.Close() // Ensure the resource is closed
        // ... process the data ...
        return nil
    }

    buffer := bytes.NewBufferString("some data")
    readCloser := io.NopCloser(buffer) // Wrap the buffer with NopCloser to satisfy io.ReadCloser

    err := processData(readCloser)
    ```

13. `io.Discard` - similar to `/dev/null` in linux. Useful for benchmarking or silencing unwanted output.
    ```go
    io.Copy(io.Discard, strings.NewReader("ignored data"))
    ```

14. `io.SectionReader` - reader for a file slice. Used to read a specific section of a file as a standalone Reader. It is useful in binary formats, partial file transfers.  
    Here is an example:
    ```go
    f, _ := os.Open("data.bin")
    section := io.NewSectionReader(f, 100, 50) // offset 100, length 50
    io.Copy(os.Stdout, section)
    ```

15. `io.ReadSeekCloser` - There are a number of interface compositions like this:  
    ```go
    type ReadSeekCloser interface {
        io.Reader
        io.Seeker
        io.Closer
    }
    ```
    Why we need it: design APIs that require multiple capabilities (e.g. file-like objects in HTTP, compression, encryption).

16. `io.WriteString(w, text)` - efficient string write, avoids byte-slice conversion.  
    ```go
    io.WriteString(os.Stdout, "fast write\n")
    ```

17. Error Sentinels: Used to control the read flow.  
    - `io.EOF` - normal end of input.
    - `io.ErrUnexpectedEOF`- stream ended too early (truncated).

18. Composability. The core reason this package exists is you can plug any `io.Reader` into any `io.Writer` regardless of what they represent:
    - File → gzip → network
    - Memory → encrypt → file
    - TCP → proxy → stdout


---

### bufio

The bufio package provides buffered I/O — wrappers around the `io.Reader` and `io.Writer` interfaces that reduce system call overhead, improve performance, and add convenience methods.

1. `bufio.NewReader(r io.Reader)` - wraps an `io.Reader` with an internal buffer (default 4KB). Reads large chunks from the underlying source, then serves your smaller reads from memory. Here is a clear example showing why `bufio` is better:  
    ```go
    // Without buffering
    file, _ := os.Open("data.txt")
    buf := make([]byte, 1)
    for {
        n, err := file.Read(buf)  // syscall per byte
        if err == io.EOF { break }
        fmt.Print(string(buf[:n]))
    }

    // --------------------------------
    // With buffering
    file, _ := os.Open("data.txt")
    reader := bufio.NewReader(file)
    for {
        b, err := reader.ReadByte() // served from buffer
        if err == io.EOF { break }
        fmt.Print(string(b))
    }
    ```
    This example clearly shows avoidance of calling system-calls in every tiny read, instead serving from a buffer.

2. `ReadByte()` / `ReadRune()` - These read one byte or one UTF-8 rune, respectively, from the internal buffer. For example:
    ```go
    r := bufio.NewReader(strings.NewReader("GÖ"))
    b, _ := r.ReadByte() // 'G'
    runeVal, size, _ := r.ReadRune() // 'Ö', 2 bytes
    fmt.Println(b, runeVal, size)
    ```
    Why: precise character-by-character reading (useful in parsers, protocols).
    Manual io.Reader handling of runes would require manual UTF-8 decoding.

3. `ReadString(delim byte)` / `ReadBytes(delim byte)` - Reads until the specified delimiter. Without `bufio` you'd have to read chunks, find delimiter manually, merge slices.  
    ```go
    r := bufio.NewReader(strings.NewReader("cmd1\ncmd2\n"))
    line, _ := r.ReadString('\n')
    fmt.Print(line) // "cmd1\n"
    ```

4. `Readline()` - Low-level method for reading long lines safely (avoids allocation explosion).  
    Returns a slice of the current line and a flag indicating continuation if line too long.
    ```go
    r := bufio.NewReader(strings.NewReader("long line here\n"))
    line, _, _ := r.ReadLine()
    fmt.Println(string(line))
    ```

5. `Peek()` - Look ahead without consuming bytes.  
    Without buffering: impossible unless you manually re-read or store bytes.  
    ```go
    r := bufio.NewReader(strings.NewReader("GET / HTTP/1.1\r\n"))
    peek, _ := r.Peek(3)
    fmt.Println(string(peek)) // "GET"
    ```

6. `Buffered()` - Returns how many bytes remain unread in the internal buffer.  
    Why: for flow control or debugging buffered state.  
    ```go
    r := bufio.NewReader(strings.NewReader("12345"))
    r.ReadByte()
    fmt.Println(r.Buffered()) // 4
    ```

7. `Discard(n)` - Skip N bytes efficiently (without allocating a slice). For example:
    ```go
    r := bufio.NewReader(strings.NewReader("abcdef"))
    r.Discard(3)
    b, _ := r.ReadByte()
    fmt.Println(string(b)) // 'd'
    ```

8. `bufio.NewWriter(w io.Writer)` - writes data into memory first, flushes to underlying writer in chunks.  
    Without buffering, we would waste syscalls:
    ```go
    for _, ch := range "hello" {
        os.Stdout.Write([]byte{byte(ch)}) // syscall per char
    }
    ```
    With buffering:
    ```go
    w := bufio.NewWriter(os.Stdout)
    for _, ch := range "hello" {
        w.WriteByte(byte(ch))
    }
    w.Flush()
    ```
    Drastically reduces the number of syscalls, essential in network servers and file I/O.

9. `Flush()` - Pushes pending data in buffer to the underlying writer.  
    Why: buffered writer doesn't automatically write until buffer full; flush ensures delivery.
    ```go
    w := bufio.NewWriter(os.Stdout)
    w.WriteString("Hello")
    w.Flush() // without this, data may remain in memory
    ```

10. `Available()` and `Buffered()` writer side  
    `Available()` → remaining space before flush needed.
    `Buffered()` → bytes currently waiting in buffer.

    ```go
    w := bufio.NewWriterSize(os.Stdout, 10)
    w.WriteString("12345")
    fmt.Println(w.Buffered(), w.Available()) // 5, 5
    ```

11. `bufio.NewReadWriter(r *io.Reader, w *io.Writer)` - Combines both reader and writer for duplex streams (TCP connections).  
    Here is an example:
    ```go
    conn, _ := net.Dial("tcp", "example.com:80")
    rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
    rw.WriteString("GET / HTTP/1.1\r\n\r\n")
    rw.Flush()
    resp, _ := rw.ReadString('\n')
    fmt.Println(resp)
    ```

12. `bufio.Scanner`
    High-level reader for token-based iteration (lines, words, custom tokens). Here is an example:
    ```go
    scanner := bufio.NewScanner(strings.NewReader("word1 word2 word3"))
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    ```

    12.1. Custom Split function for `Scanner`   
    We can define our own tokenization logic as anonymous function inside `scanner.Scan()`

13. `NewReaderSize` and `NewWriterSize` - Allows specifying custom buffer capacity.   
    Why: tune memory vs. syscall tradeoff depending on I/O pattern.
    ```go
    r := bufio.NewReaderSize(os.Stdin, 64*1024)
    ```
    This example shows large buffer for big sequential reads (for example file copy)

14. `Reset()` methods - Reuses a buffer with a new underlying source/sink.  
    Why: avoid allocating new readers/writers repeatedly in tight loops.
    ```go
    r := bufio.NewReader(strings.NewReader("old"))
    r.Reset(strings.NewReader("new"))
    line, _ := r.ReadString('\n')
    fmt.Println(line)
    ```

15. `bufio` in networking context  
    TCP reads/writes are non-deterministic — may return partial packets.  
    `bufio` handles:
    - Fragmented reads.
    - Reduced syscalls.
    - Framing convenience (ReadLine, Peek).
    
    ```go
    listener, _ := net.Listen("tcp", ":9000")
    for {
        conn, _ := listener.Accept()
        go func(c net.Conn) {
            r := bufio.NewReader(c)
            msg, _ := r.ReadString('\n')
            fmt.Println("Received:", msg)
            c.Close()
        }(conn)
    }
    ```