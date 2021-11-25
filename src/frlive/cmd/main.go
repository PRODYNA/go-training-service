package main

import (
	"github.com/gorilla/mux"
	"github.com/prodyna/frlive/adapter"
	"github.com/prodyna/frlive/config"
	"github.com/prodyna/frlive/port"
	"github.com/prodyna/frlive/service"
	"github.com/prodyna/goprobes/probes"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	Service *service.ServiceImpl
	Probes  *mux.Router
	Ports   *mux.Router
	Config  *config.Config
}

func main() {
	s := NewServer()
	err := s.Configure()
	s.Init()
	s.InitProbes(err)
	s.InitPorts()

	go func() {
		port := s.Config.GetProbePort()
		log.Err(http.ListenAndServe(port, s.Probes)).Msg("cant listen to probes router")
	}()

	go func() {
		port := s.Config.GetPort()
		log.Info().Str("port", port).Msg("start listening to port")
		log.Err(http.ListenAndServe(port, s.Ports)).Msg("cant listen to ports")
	}()

	waitForSignal()
}

func NewServer() *server {
	log.Info().Msg("creating new server")

	adapter := adapter.NewAdapter(http.DefaultClient)

	return &server{
		Service: service.NewService(http.DefaultClient, adapter),
		Probes:  &mux.Router{},
		Ports:   &mux.Router{},
	}
}

func (s server) Init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("starting server")
}

func (s server) InitProbes(configError error) {
	log.Info().Msg("starting probes")
	ps := probes.NewProbeService()
	ps.HandleProbes(s.Probes)

	ps.AddStart(func() (string, bool) {
		if configError == nil {
			return "ok", true
		} else {
			return configError.Error(), false
		}
	})
}

func (s server) InitPorts() {

	userPort := port.NewUserPort(s.Service)
	userPort.HandleHttp(s.Ports)
}

func (s *server) Configure() error {

	c, err := config.LoadConfig()
	s.Config = c

	if err != nil {
		return err
	}

	return nil
}

func waitForSignal() {
	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)

	signal := <-s

	log.Info().Str("sig", signal.String()).Msg("Terminating")
}
