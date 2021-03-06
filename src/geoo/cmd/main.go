package main

import (
	"github.com/gorilla/mux"
	"github.com/prodyna/go-training/geoo/adapter"
	"github.com/prodyna/go-training/geoo/config"
	"github.com/prodyna/go-training/geoo/port"
	c "github.com/prodyna/goconfig/config"
	"github.com/prodyna/goprobes/probes"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)


type server struct{
	Config config.Configuration
}

func main() {

	log.Info().Msg("starting server")

	config := config.Configuration{}
	err := c.NewConfigLoader().LoadConfig(&config)
	if err != nil {
		log.Err(err).Msg("cannot load config")
	}

	s := server{
		Config: config,
	}

	s.Init()
	go s.InitProbes()
	go s.InitMetrics()
	go s.InitWeb()
	s.WaitForTerminate()
}

func (s server) InitWeb() {
	log.Info().Str("port" , s.Config.Web.GetPort()).Msg("starting web")
	router := mux.NewRouter()

	a := adapter.NewHttpBin(s.Config, http.DefaultClient)
	port.HandleUserPort(router, a)

	log.Err(http.ListenAndServe(s.Config.Web.GetPort(), router))
}

func (s server) Init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func (s server) InitProbes() {
	router := mux.NewRouter()
	ps := probes.NewProbeService()
	ps.HandleProbes(router)
	log.Err(http.ListenAndServe(s.Config.Probe.GetPort(), router))
}

func (s server) WaitForTerminate()  {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
	signal := <- sig
	log.Info().Str("signal", signal.String()).Msg("terminating")
}

func (s server) InitMetrics() {
	log.Info().Str("port" , s.Config.Metrics.GetPort()).Msg("starting metrics")
	metrics := promhttp.Handler()
	log.Err(http.ListenAndServe(s.Config.Metrics.GetPort(), metrics))
}