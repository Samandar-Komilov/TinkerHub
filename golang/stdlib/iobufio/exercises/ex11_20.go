package exercises

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Ex11() {
	fp := filepath.Join(BASE_DIR, "files", "in.txt")
	r, _ := os.Open(fp)
	defer r.Close()

	res, _ := io.ReadAll(r)
	fmt.Println(string(res))
}

type TimeStampedWriter struct {
	w io.Writer
}

func (tw *TimeStampedWriter) Write(p []byte) (n int, err error) {
	timestamp := time.Now().Format(time.RFC3339)
	lines := strings.Split(string(p), "\n")
	for i := range lines {
		if lines[i] != "" {
			lines[i] = string(timestamp) + " " + lines[i]
		}
	}
	res := strings.Join(lines, "\n")

	return tw.w.Write([]byte(res))
}

func Ex12() {
	fp := filepath.Join(BASE_DIR, "files", "hello.txt")
	f, _ := os.Create(fp)
	defer f.Close()

	writer := &TimeStampedWriter{w: f}
	writer.Write([]byte("Hello World!\nWhats up?\n"))

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

func Ex13() {
	fpath := filepath.Join(BASE_DIR, "files", "in.txt")
	large_file, err := os.Open(fpath)
	if err != nil {
		log.Fatal("Error while opening file:", err)
	}
	defer large_file.Close()

	dest_file, err := os.Create("files/copied.log")
	if err != nil {
		log.Fatal("Error while creating destination file:", err)
	}
	defer dest_file.Close()

	io.Copy(dest_file, large_file)
}

func Ex14() {
	fpath := filepath.Join(BASE_DIR, "files", "in.txt")
	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fr := io.LimitReader(file, 10)

	dest_file, err := os.Create("files/copied.log")
	if err != nil {
		log.Fatal(err)
	}
	defer dest_file.Close()

	io.Copy(dest_file, fr)
}

func Ex15() {
	f1, _ := os.Open(filepath.Join(BASE_DIR, "files", "in2.txt"))
	f2, _ := os.Open(filepath.Join(BASE_DIR, "files", "in3.txt"))
	defer f1.Close()
	defer f2.Close()

	mr := io.MultiReader(f1, f2)

	io.Copy(os.Stdout, mr)
}

func Ex16() {
	f1, _ := os.Open(filepath.Join(BASE_DIR, "files", "in2.txt"))
	logFile, _ := os.Create(filepath.Join(BASE_DIR, "files", "logs.log"))
	defer f1.Close()
	defer logFile.Close()

	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		n, _ := io.Copy(os.Stdout, f1) // write file contents into pipe
		fmt.Fprint(pw, n)
	}()

	io.Copy(logFile, pr) // read from pipe into log
}

func Ex17() {
	var src, dest string

	flag.StringVar(&src, "src", "", "Source file")
	flag.StringVar(&dest, "dest", "", "Destination")

	flag.Parse()

	fmt.Println(src, dest)
	// curdir, _ := os.Getwd()
	// src_path, dest_path := filepath.Join(curdir, "files", src), filepath.Join(curdir, "files", dest)

	fsrc, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	fdest, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer fsrc.Close()
	defer fdest.Close()

	written, _ := io.Copy(fdest, fsrc)
	fmt.Println("Written:", written, "bytes")
}

// func Ex18() {
// 	var file_src string
// 	flag.StringVar(&file_src, "src", "", "Source file")

// 	flag.Parse()

// }

// type SkipReader struct{}

// func (sr *SkipReader) Read(p []byte) (n int64, err error) {

// }
