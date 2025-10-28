package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Ex1str() {
	s := "This is Uzb"
	fmt.Println(strings.Contains(s, "Uzb"))
	fmt.Println(strings.Index(s, "is"))
}

func Ex2str() {
	s := "Hello $, what is your $\n"
	cnt := strings.Count(s, "$")
	s = strings.ReplaceAll(s, "$", "name")
	fmt.Printf("String: %s\nCount: %d\n", s, cnt)
}

func Ex3str() {
	s := "word1,word2,word3"
	sl := strings.Split(s, ",")
	sr := strings.Join(sl, "-")

	fmt.Println(sr)
}

func Ex5str() {
	s := "easy string"
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.ToTitle(s))
}

func Ex6str() {
	var sb strings.Builder

	sb.WriteString("Hello")
	sb.WriteByte('\t')
	sb.WriteString("World!")

	fmt.Println(sb.String())
}

func Ex7str() {
	s := "English Grammar In Use"
	fmt.Println(len(strings.Split(s, " ")))
	fmt.Println(len(s))

	cnt := 0
	vowels := "aeiuoAEIUO"
	for _, c := range s {
		for _, v := range vowels {
			if c == v {
				cnt += 1
			}
		}
	}
	fmt.Println("Num of vowels:", cnt)
}

func Ex8str() {
	s := "Great China Wall"
	res := strings.Map(func(r rune) rune {
		return r + 3
	}, s)
	fmt.Println("Result:", res)
}

func Ex9str() {
	s := "Po and panda"
	before, after, is_found := strings.Cut(s, "and")

	fmt.Println(before, after, is_found)
}

func Ex10str() {
	curdir, _ := os.Getwd()
	fpath := filepath.Join(curdir, "files", "test.csv")
	f, _ := os.Open(fpath)
	defer f.Close()

	filedata, _ := io.ReadAll(f)

	slc := strings.FieldsFunc(string(filedata), func(r rune) bool {
		return r == ','
	})

	fmt.Println(slc)
}

func Ex14str() {
	s := "Hello Eshmat!"
	rslice := []rune(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rslice[i], rslice[j] = rslice[j], rslice[i]
	}

	fmt.Println(string(rslice))
}
