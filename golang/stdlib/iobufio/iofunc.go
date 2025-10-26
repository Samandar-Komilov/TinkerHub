package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReaderExample() {
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

func WriterExample() {
	f, _ := os.Create("hello.txt")
	defer f.Close()
	f.Write([]byte("Hello World!"))
	fmt.Println("File write completed!")
}

func CopyExample() {
	src, _ := os.Open("files/in.txt")
	dst, _ := os.Create("files/out.txt")
	defer src.Close()
	defer dst.Close()
	io.Copy(dst, src)
}

func CopyNExample() {
	src, _ := os.Open("files/in.txt")
	dst, _ := os.Create("files/out.txt")
	defer src.Close()
	defer dst.Close()
	io.CopyN(dst, src, 10)
}

func CopyBufferExample() {
	src, _ := os.Open("files/in.txt")
	dst, _ := os.Create("files/out.txt")
	defer src.Close()
	defer dst.Close()
	buffer := make([]byte, 5)
	io.CopyBuffer(dst, src, buffer)
}

func LimitReaderExample() {
	r := strings.NewReader("ABCDEFGHIJK")
	limited := io.LimitReader(r, 5)
	io.Copy(os.Stdout, limited) // prints ABCDE

	fmt.Println("\nDone!")
}

func MultiReaderExample() {
	r := io.MultiReader(strings.NewReader("Hello"), strings.NewReader("World!"))
	io.Copy(os.Stdout, r)
}

func MultiWriterExample() {
	f, _ := os.Create("files/hello.txt")
	mw := io.MultiWriter(os.Stdout, f)
	mw.Write([]byte("Hello")) // prints and logs

}

func TeeReaderExample() {
	src := strings.NewReader("secret data")
	log := &bytes.Buffer{}
	tee := io.TeeReader(src, log)
	io.Copy(io.Discard, tee)
	fmt.Println("Logged:", log.String(), log)
}

func ReadAllExample() {
	r := strings.NewReader("all at once")
	data, _ := io.ReadAll(r)
	fmt.Println(string(data))
}

func PipeExample() {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "streamed data")
		w.Close()
	}()

	io.Copy(os.Stdout, r)
}

func NopCloserExample() {
	r := strings.NewReader("test")
	rc := io.NopCloser(r)
	io.Copy(os.Stdout, rc)
	defer rc.Close()
}

func DiscardExample() {
	r := strings.NewReader("test")
	io.Copy(io.Discard, r)
}

func SectionReaderExample() {
	f, _ := os.Open("data.bin")
	section := io.NewSectionReader(f, 100, 50) // offset 100, length 50
	io.Copy(os.Stdout, section)
}
