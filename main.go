package main

import (
"net/http"
"time"
"github.com/go-chi/chi"
"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(15 * time.Second) 
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}
