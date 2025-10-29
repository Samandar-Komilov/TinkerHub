package exercises

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func Ex21_TCP_echo_server() {
	listener, err := net.Listen("tcp", ":8003")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	log.Println("TCP server is listening at :8003...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			tc := io.TeeReader(c, os.Stdout)
			io.Copy(c, tc)
			c.Close()
		}(conn)
	}
}

func Ex22_TCP_client() {
	conn, err := net.Dial("tcp", ":8003")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var input string
	fmt.Scanf("%s", &input)
	r := strings.NewReader(input)

	written, err := io.Copy(conn, r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sent", written, "bytes to server.")
}

func Ex23_simple_proxy() {
	proxyAddr := ":8003"
	proxyListener, err := net.Listen("tcp", proxyAddr)
	if err != nil {
		panic(err)
	}
	defer proxyListener.Close()

	for {
		clientConn, err := proxyListener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		go handleConn(clientConn)
	}
}

func handleConn(clientConn net.Conn) {
	defer clientConn.Close()

	buf := make([]byte, 1024)
	_, err := clientConn.Read(buf)
	if err != nil {
		log.Fatal(err)
		return
	}

	upstreamAddr := ":8004"
	upstreamConn, err := net.Dial("tcp", upstreamAddr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer upstreamConn.Close()

	_, err = upstreamConn.Write(buf)
	if err != nil {
		log.Fatal(err)
		return
	}

	go io.Copy(clientConn, upstreamConn)
}

func Ex24_file_server() {
	ln, err := net.Listen("tcp", ":8002")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Println("Server listening on port 8002...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Connection accept error:", err)
			continue
		}

		go handleConnFileServe(conn)
	}
}

func handleConnFileServe(conn net.Conn) {
	defer conn.Close()

	// ✅ Read client request first
	reader := bufio.NewReader(conn)
	request, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Failed to read request:", err)
		conn.Write([]byte("ERROR: Failed to read request\n"))
		return
	}

	// ✅ Parse filename from request (simple protocol: "GET filename\n")
	request = strings.TrimSpace(request)
	if !strings.HasPrefix(request, "GET ") {
		conn.Write([]byte("ERROR: Invalid request format. Use: GET filename\n"))
		return
	}

	filename := strings.TrimPrefix(request, "GET ")
	if filename == "" {
		conn.Write([]byte("ERROR: No filename specified\n"))
		return
	}

	// ✅ Security: Prevent directory traversal
	if strings.Contains(filename, "..") || strings.Contains(filename, "/") {
		conn.Write([]byte("ERROR: Invalid filename\n"))
		return
	}

	// ✅ Construct safe file path
	curdir, _ := os.Getwd()
	filePath := filepath.Join(curdir, "files", filename)

	// ✅ Verify the file exists and is within allowed directory
	if !strings.HasPrefix(filePath, filepath.Join(curdir, "files")) {
		conn.Write([]byte("ERROR: Access denied\n"))
		return
	}

	// ✅ Open the requested file
	f, err := os.Open(filePath)
	if err != nil {
		log.Println("File open error:", err)
		conn.Write([]byte("ERROR: File not found or cannot be opened\n"))
		return
	}
	defer f.Close()

	// ✅ Get file info for basic headers
	fileInfo, err := f.Stat()
	if err != nil {
		log.Println("File stat error:", err)
		conn.Write([]byte("ERROR: Cannot access file information\n"))
		return
	}

	// ✅ Send success response with file size info
	conn.Write([]byte(fmt.Sprintf("OK %d\n", fileInfo.Size())))

	// ✅ Send file content
	_, err = io.Copy(conn, f)
	if err != nil {
		log.Println("File copy error:", err)
	}
}
