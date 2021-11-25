package main

import (
	"github.com/gorilla/mux"
	"github.com/prodyna/fratschin/adapter"
	"github.com/prodyna/fratschin/cfg"
	"github.com/prodyna/fratschin/port"
	"github.com/prodyna/goprobes/probes"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

type server struct{}

func main() {

	config := cfg.Load()
	s := server{}
	s.InitLogging()
	log.Info().Msg("starting server")

	go s.InitProbes(config)
	go s.InitMetrics(config)
	go s.InitRest(config)

	s.WaitForTerminate()

}

func (s server) InitMetrics(cfg *cfg.Config) {
	log.Info().Str("port", cfg.Metrics.Port).Msg("starting metrics")
	metric := promhttp.Handler()
	log.Err(http.ListenAndServe(":"+cfg.Metrics.Port, metric))
}

func (s server) InitProbes(cfg *cfg.Config) {
	router := mux.NewRouter()
	ps := probes.NewProbeService()
	ps.HandleProbes(router)

	log.Info().Str("port", cfg.Probe.Port).Msg("starting probes")
	http.ListenAndServe(":"+cfg.Probe.Port, router)
}

func (s server) InitRest(cfg *cfg.Config) {
	router := mux.NewRouter()
	a := adapter.NewAdapter(http.DefaultClient)
	port.HandleRest(router, a)
	log.Info().Str("port", cfg.Rest.Port).Msg("starting rest")
	http.ListenAndServe(":"+cfg.Rest.Port, router)
}

func (s server) InitLogging() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func (s server) WaitForTerminate() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
	signal := <-sc
	log.Info().Str("signal", signal.String()).Msg("terminating")
}
