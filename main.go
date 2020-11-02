package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"os"
	"strconv"
)

var cont int64 = 1

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome - server name: "+os.Getenv("SERVER")+" - request number: "+strconv.FormatInt(cont, 10)))
		cont++
	})
	http.ListenAndServe(":3000", r)
}

