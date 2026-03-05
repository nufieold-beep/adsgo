
package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

type BidderConfig struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Endpoint string `yaml:"endpoint"`
	Timeout int `yaml:"timeout"`
}

type Config struct {
	Bidders []BidderConfig `yaml:"bidders"`
}

func Load(path string) (*Config,error) {

	data,err := os.ReadFile(path)

	if err != nil {
		return nil,err
	}

	var cfg Config

	err = yaml.Unmarshal(data,&cfg)

	return &cfg,err
}
