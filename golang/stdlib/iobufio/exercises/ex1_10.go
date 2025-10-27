package exercises

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	BASE_DIR = "/home/voidp/Projects/oss/TinkerHub/golang/stdlib/iobufio"
)

func Ex1() {
	rdr := strings.NewReader("Hello World!\n")

	io.Copy(os.Stdout, rdr)
}

type HelloReturn struct{}

func (h HelloReturn) Read() string {
	return "Hello"
}

func Ex2() {
	h := HelloReturn{}
	fmt.Println(h.Read())
}

type LogWriter struct{}

func (l LogWriter) Write(s string) {
	fmt.Fprintf(os.Stdout, "[LOG] %s\n", s)
}

func Ex3() {
	l1 := LogWriter{}
	l1.Write("This is a test")
}

func Ex4() {
	rdr := strings.NewReader("Hello my dear.\n")
	io.CopyN(os.Stdout, rdr, 10)
}

func Ex5() {
	fp := filepath.Join(BASE_DIR, "files", "in.txt")
	f, err := os.Open(fp)
	if err != nil {
		log.Fatal("Error while opening a file:", err)
	}

	lr := io.LimitReader(f, 18)

	io.Copy(os.Stdout, lr)
}

func Ex6() {
	r1 := strings.NewReader("First string\n")
	r2 := strings.NewReader("Second String\n")

	r := io.MultiReader(r1, r2)

	io.Copy(os.Stdout, r)
}

func Ex7_8() {
	fp := filepath.Join(BASE_DIR, "files", "write2.txt")
	w2, err := os.Create(fp)
	if err != nil {
		log.Fatal("Error while creating file:", err)
	}
	defer w2.Close()

	w := io.MultiWriter(os.Stdout, w2)

	r := strings.NewReader("This is another string.\n")
	n, _ := io.Copy(w, r)

	fmt.Println("Copied data:", n, "bytes")
}

func Ex9() {
	fp := filepath.Join(BASE_DIR, "files", "write2.txt")
	f, _ := os.Open(fp)
	defer f.Close()

	logp := filepath.Join(BASE_DIR, "files", "logs.txt")
	logf, _ := os.Open(logp)
	defer logf.Close()

	r := io.TeeReader(f, os.Stdout)
	io.Copy(logf, r)
}

func Ex10() {
	r1 := strings.NewReader("String 1")
	r2 := strings.NewReader("String 1")

	s1, _ := io.ReadAll(r1)
	s2, _ := io.ReadAll(r2)

	is_equal := bytes.Compare(s1, s2)
	if is_equal == 0 {
		fmt.Println("Equal stream data!")
	} else {
		fmt.Println("Not Equal stream data!")
	}
}
