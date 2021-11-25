package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
)

type Config struct {
	Port      int `yaml:"port"`
	ProbePort int `yaml:"probePort"`
}

func (c Config) GetProbePort() string {
	if c.Port == 0 {
		return ":8081"
	}
	return ":" + strconv.Itoa(c.ProbePort)
}

func (c Config) GetPort() string {
	return ":" + strconv.Itoa(c.Port)
}

func LoadConfig() (*Config, error) {
	c := &Config{}
	data, err := os.ReadFile("config.yaml")

	if err != nil {
		return c, err
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		return c, err
	}

	return c, nil

}
