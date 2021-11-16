package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

type server struct{}

func main() {
	s := server{}

	log.Printf("Starting server")

	go func() {
		log.Fatal(http.ListenAndServe(":8083",promhttp.Handler()))
	}()

	log.Fatal(http.ListenAndServe(":8080", &s))
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`{ "ok" : true , "status": "ok", "instance": "fratschi", "version": "5"}`))
	log.Printf("One request from %s", r.RemoteAddr)
}
