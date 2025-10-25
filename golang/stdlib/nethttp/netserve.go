package main

import (
	"log"
	"net/http"
)

func Main_netsrv() {
	log.Println("Listening on port 8000...")

	srv := &http.Server{
		Addr: ":8001",
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
}
