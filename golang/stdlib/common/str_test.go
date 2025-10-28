package main

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkStringBuilder(b *testing.B) {
	builder := strings.Builder{}
	for i := 0; i < b.N; i++ {
		builder.WriteString("Hello, ")
		builder.WriteString("world!")
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	buffer := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		buffer.WriteString("Hello, ")
		buffer.WriteString("world!")
	}
}

/*
Run: go test -bench=. -benchmem=0
*/
