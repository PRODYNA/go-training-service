package main

import (
	"github.com/gorilla/mux"
	"github.com/prodyna/fratschin/adapter"
	cfg2 "github.com/prodyna/fratschin/cfg"
	"github.com/prodyna/fratschin/port"
	"github.com/prodyna/goconfig/config"
	"github.com/prodyna/goprobes/probes"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

type server struct {}


func main() {

	cfg := &cfg2.Config{}
	err := config.NewConfigLoader().LoadConfig(cfg)
	if err != nil {
		log.Info().Msg("cannot load config")

	}
	s := server{}
	s.InitLogging()
	log.Info().Msg("starting server")

	go s.InitProbes(cfg)
	go s.InitMetrics()
	go s.InitRest(cfg)

	s.WaitForTerminate()

}


func (s server) InitMetrics() {
	log.Info().Str("port", ":8083").Msg("starting metrics")
	metric := promhttp.Handler()
	log.Err(http.ListenAndServe(":8083", metric))
}

func (s server) InitProbes(cfg *cfg2.Config) {
	router := mux.NewRouter()
	ps := probes.NewProbeService()
	ps.HandleProbes(router)

	log.Info().Str("port", cfg.Probe.Port).Msg("starting probes")
	http.ListenAndServe(":8081", router)
}

func (s server) InitRest(cfg *cfg2.Config) {
	router := mux.NewRouter()
	a := adapter.NewAdapter(http.DefaultClient)
	port.HandleRest(router, a)
	log.Info().Str("port", cfg.Probe.Port).Msg("starting rest")
	http.ListenAndServe(":8080", router)
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
