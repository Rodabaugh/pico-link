package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	port := "8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/admin/", func(w http.ResponseWriter, r *http.Request) {
		MainPage().Render(r.Context(), w)
	})
	mux.HandleFunc("GET /google", handlerRedirect)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting pico-link on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
