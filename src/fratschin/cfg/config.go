package cfg


type Config struct {
	Rest  Rest  `yaml:"rest"`
	Probe Probe `yaml:"probe"`
}
type Rest struct {
	Port string `yaml:"port"`
}
type Probe struct {
	Port string `yaml:"port"`
}



