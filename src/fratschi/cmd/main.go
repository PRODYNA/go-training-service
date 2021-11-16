package main

import (
	"github.com/gorilla/mux"
	"github.com/prodyna/goprobes/probes"
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
	s.WaitForTerminate()

}

func (s server) InitProbes() {
	router := mux.NewRouter()
	ps := probes.NewProbeService()
	ps.HandleProbes(router)
	log.Err(http.ListenAndServe(":8081", router))
}

func (s server) Init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func (s server) WaitForTerminate() {

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
	receivedSignal := <-sig
	log.Info().Str("signal", receivedSignal.String()).Msg("terminating")
}
