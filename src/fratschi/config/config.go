package config

type Configuration struct {
	Web     Web     `yaml:"web"`
	Probe   Probe   `yaml:"probe"`
	Metrics Metrics `yaml:"metrics"`
}

type Web struct {
	Port string `yaml:"port"`
}

type Probe struct {
	Port string `yaml:"port"`
}

type Metrics struct {
	Port string `yaml:"port"`
}

func (m Metrics) GetPort() string {
	return ":" + m.Port
}

func (m Probe) GetPort() string {
	return ":" + m.Port
}

func (m Web) GetPort() string {
	return ":" + m.Port
}
