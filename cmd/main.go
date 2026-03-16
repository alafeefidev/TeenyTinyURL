package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var urls = map[string]string{
	"IjhYW": "https://google.com",
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello to hmd.sh")
	})

	r.Get("/{short_id}", func(w http.ResponseWriter, r *http.Request) {
		shortID := chi.URLParam(r, "short_id")
		url, ok := urls[shortID]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		http.Redirect(w, r, url, http.StatusPermanentRedirect)
	})

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "healthy")
	})

	fmt.Println("Server starting at :6767")
	err := http.ListenAndServe(":6767", r)
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
