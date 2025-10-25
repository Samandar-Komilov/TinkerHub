package main

import (
	"log"
	"net/http"
	"time"
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("before request handled:", time.Now())
		next.ServeHTTP(w, r)
		log.Println("after request handled:", time.Now())
	})
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, "executing fooHandler")
	w.Write([]byte("OK"))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, "executing barHandler")
	w.Write([]byte("OK"))
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /foo", http.HandlerFunc(fooHandler))
	mux.Handle("GET /bar", http.HandlerFunc(barHandler))

	log.Print("listening on :8090...")
	err := http.ListenAndServe(":8090", logMiddleware(mux))
	log.Fatal(err)
}
