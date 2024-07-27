package main

import (
	"log"
	"net/http"
)

func main() {
	const filePathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(filePathRoot)))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving %s on port: %s\n", filePathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
