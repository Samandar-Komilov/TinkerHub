package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
	count int
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.count++
	fmt.Fprintf(w, "Request Count: %d", h.count)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello\n")
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v:%v\n", name, h)
		}
		fmt.Println(name, headers)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.Handle("/", &MyHandler{})

	fmt.Println("Listening on port 8090...")
	http.ListenAndServe(":8090", nil)
}
