package cfg

import (
	"github.com/prodyna/goconfig/config"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Rest    Port `yaml:"rest"`
	Probe   Port `yaml:"probe"`
	Metrics Port `yaml:"metrics"`
}

type Port struct {
	Port string `yaml:"port"`
}

func (p Port) GetPort() string {
	return ":" + p.Port
}

func Load() *Config {

	cfg := &Config{}
	err := config.NewConfigLoader().LoadConfig(cfg)
	if err != nil {
		log.Info().Msg("cannot load config")
	}
	return cfg
}
