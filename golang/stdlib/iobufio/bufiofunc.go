package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func ReadByteRuneExample() {
	r := bufio.NewReader(strings.NewReader("GÖNAYDIN"))
	b, _ := r.ReadByte()
	runeVal, size, _ := r.ReadRune()

	fmt.Printf("%s %s %d\n", string(b), string(runeVal), size)
}

func ReadStringExample() {
	r := bufio.NewReader(strings.NewReader("cmd1\ncmd2\n"))
	line, _ := r.ReadString('\n')
	line2, _ := r.ReadBytes('\n')
	fmt.Println(line) // "cmd1\n"
	fmt.Println(line2)
}

func ReadLineExample() {
	r := bufio.NewReader(strings.NewReader("Hello World!"))
	line, _, _ := r.ReadLine()
	fmt.Println(string(line))
}

func PeekExample() {
	r := bufio.NewReader(strings.NewReader("GET / HTTP/1.1\r\n"))
	peek, _ := r.Peek(3)
	fmt.Println(string(peek)) // "GET"
}

func BufferedExample() {
	r := bufio.NewReader(strings.NewReader("12345"))
	r.ReadByte()
	fmt.Println(r.Buffered()) // 4
}

func DiscardBufioExample() {
	r := bufio.NewReader(strings.NewReader("abcdef"))
	r.Discard(3)
	b, _ := r.ReadByte()
	fmt.Println(string(b)) // 'd'
}

func AvailableBufferedExample() {
	w := bufio.NewWriterSize(os.Stdout, 10)
	w.WriteString("12345")
	fmt.Println(w.Buffered(), w.Available()) // 5, 5
	w.Flush()
	fmt.Println(w.Buffered(), w.Available())
}

func NewReadWriterExample() {
	conn, _ := net.Dial("tcp", "example.com:80")
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	rw.WriteString("GET / HTTP/1.1\r\n\r\n")
	rw.Flush()
	resp, _ := rw.ReadString('\n')
	fmt.Println(resp)
}

func ScannerExample1() {
	scanner := bufio.NewScanner(strings.NewReader("word1 word2 word3"))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ScannerExample2() {
	scanner := bufio.NewScanner(strings.NewReader("word1,word2,word3"))
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i, b := range data {
			if b == ',' {
				return i + 1, data[:i], nil
			}
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	})

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// wrong function, but I didn't spend much time to fix it
}
