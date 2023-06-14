package conf

import (
	_ "embed"
	"os"

	"gopkg.in/yaml.v3"
)

// go: embed config.yaml
var s []byte

const (
	Debug = "debug"
	Dev   = "dev"
)

type Config struct {
	Name string
	Mode string
	Log  struct {
		Level string
	}
}

// ParseConf express parse config.yaml to Config struct
func ParseConf(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c Config
	return &c, yaml.Unmarshal(b, &c)
}
