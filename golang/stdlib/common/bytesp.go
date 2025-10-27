package main

import (
	"bufio"
	"bytes"
	"fmt"
	"time"
)

func Ex1() {
	byteSlice := []byte("!Hello World$")
	fmt.Println(bytes.HasPrefix(byteSlice, []byte("!")))
	fmt.Println(bytes.HasSuffix(byteSlice, []byte("$")))
}

func Ex2() {
	data := []byte("Hello\nThis is new line\n\nMy name is\n\tGerrard!")
	lines := bytes.SplitSeq(data, []byte("\n"))
	lines2 := bytes.Fields([]byte("Hi\nwhat up?"))

	for line := range lines {
		fmt.Println(string(line))
	}

	fmt.Println(lines2)
}

func Ex3() {
	bs1 := []byte("Byte 1")
	bs2 := []byte("Byte 2")
	bs := make([][]byte, 0, len(bs1)+len(bs2))
	bs = append(bs, bs1)
	bs = append(bs, bs2)

	fmt.Println(bytes.Join(bs, []byte(">")))
}

func Ex4() {
	bs := []byte("&ll &s h&ve been repl&ced!")
	bs = bytes.ReplaceAll(bs, []byte("&"), []byte("a"))
	fmt.Println(string(bs))
}

func Ex5() {
	bs := []byte("Donald Duck is the richest cartoon actor in the world!")
	fmt.Println(bytes.Index(bs, []byte("ck")))
	fmt.Println(bytes.LastIndex(bs, []byte("D")))
	fmt.Println(bytes.Contains(bs, []byte("actor")))
}

func Ex6() {
	bs := []byte("This Is Case Sensitive!")
	bs2 := []byte("this is case sensitive!")

	fmt.Println(bytes.EqualFold(bs, bs2))
	fmt.Println(bytes.Equal(bs, bs2))
}

func Ex7() {
	bs := []byte("   Hello World!   ")
	// bs = bytes.Trim(bs, " ")
	bs = bytes.TrimSpace(bs)
	fmt.Println(string(bs))
}

func Ex8() {
	mystr := "Toshmat"
	fmt.Println([]byte(mystr))
}

func Ex9() {
	r := bytes.NewReader([]byte("This is a bytes\n reader source.\n"))
	// read, _ := io.ReadAll(r)
	// fmt.Println(read)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func Ex11() {
	buf := &bytes.Buffer{}

	for range 3 {
		buf.WriteString("[LOG] ")
		buf.WriteString(time.Now().Format(time.RFC3339))
		buf.WriteByte('\n')
	}

	fmt.Println(buf.String(), buf.Cap(), buf.Len())

	data := make([]byte, 5)
	n, _ := buf.Read(data)
	fmt.Printf("%d bytes read: %s\n", n, string(data))
}

func Ex12() {
	buf := &bytes.Buffer{}
	buf.WriteString("Hello World!\n")
	fmt.Println(buf.String(), buf.Cap(), buf.Len())
	buf.Reset()
	fmt.Println(buf.String(), buf.Cap(), buf.Len())
	buf.Grow(100)
	fmt.Println(buf.String(), buf.Cap(), buf.Len())
	buf.WriteString("Another string!")
	buf.Truncate(4)
	fmt.Println(buf.String(), buf.Cap(), buf.Len())
}

func Ex13() {
	buf := &bytes.Buffer{}
	x := 7
	fmt.Fprintf(buf, "%d", x)
	fmt.Println(buf.String())
}
