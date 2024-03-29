package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

var cont int64 = 1

func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	exposeOnAddress := os.Getenv("EXPOSE_ON_ADDRESS")
	exposeOnPort := os.Getenv("EXPOSE_ON_PORT")
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome - local ip: " + getIP() + " - request number: " + strconv.FormatInt(cont, 10)))
		cont++
	})
	http.ListenAndServe(fmt.Sprintf("%s:%s", exposeOnAddress, exposeOnPort), r)
}
