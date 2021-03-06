package main

import (
	"github.com/gorilla/mux"
	"github.com/prodyna/go-training-service/tree/team/sherdt/web"
	"github.com/prodyna/goprobes/probes"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
}

func main() {
	log.Info().Msg("starting server")

	s := server{}
	s.Init()
	go s.InitProbes()
	go s.InitMetrics()
	go s.InitWeb()

	s.WaitForProbes()
}

func (s server) InitWeb() {
	router := mux.NewRouter()
	web.HandleWebPort(router)
	log.Info().Str("port", ":8080").Msg("starting web port")
	log.Err(http.ListenAndServe(":8080", router))
}

func (s server) Init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func (s server) InitProbes() {
	router := mux.NewRouter()
	ps := probes.NewProbeService()
	ps.HandleProbes(router)

	log.Err(http.ListenAndServe(":8081", router))
}

func (s server) InitMetrics() {
	metric := promhttp.Handler()
	log.Err(http.ListenAndServe(":8082", metric))
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, err := w.Write([]byte(`{ "ok" : true , "status": "ok", "instance": "template", "version": "4"}`))
	if err != nil {
		return
	}
	log.Printf("One request from %s", r.RemoteAddr)
}

func (s server) WaitForProbes() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
	receivedSignal := <-sig
	log.Info().Str("signal", receivedSignal.String()).Msg("terminating")
}
