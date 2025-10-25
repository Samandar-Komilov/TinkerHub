package main

import (
	"log"
	"net/http"
)

func Main_muxers() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	log.Println("Listening on port 8090...")
	err := http.ListenAndServe(":8090", mux)
	if err != nil {
		log.Fatal(err)
		return
	}
}
