package main

import (
	"log"
	"net/http"
)

type server struct {}

func main() {
  s := server{}

  log.Fatal(http.ListenAndServe(":8080", &s))
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`{ "ok" : true , "status": "ok", "instance": "template", "version": "2"}`))
	log.Printf("One request from %s", r.RemoteAddr )
}
