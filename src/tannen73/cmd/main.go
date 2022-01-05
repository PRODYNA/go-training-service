package main

import (
	"github.com/gorilla/mux"
	"github.com/prodyna/goprobes/probes"
	"log"
	"net/http"
)

type server struct{}

func main() {
	pRouter := mux.NewRouter()

	initProbes(pRouter)
	initServer(pRouter)

	log.Printf("Starting server")
	log.Fatal(http.ListenAndServe(":8080", pRouter))
}

func initProbes(pRouter *mux.Router) {
	ps := probes.NewProbeService()

	ps.AddStart(func() (string, bool) {
		return "Started", true
	})

	ps.AddLive(probes.NewMemoryProbe(1024 * 1024))

	ps.AddReady(func() (string, bool) {
		return "Ready", true
	})

	ps.HandleProbes(pRouter)
}

func initServer(pRouter *mux.Router) {
	s := server{}
	pRouter.
		Methods("GET").
		Path("/").
		Handler(s)
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`{ "ok" : true , "status": "ok", "instance": "tannen73", "version": "4"}`))
	log.Printf("One request from %s", r.RemoteAddr)
}
