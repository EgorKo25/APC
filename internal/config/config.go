package config

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/caarlos0/env/v9"
)

func NewConfig() (*Config, error) {

	c := &Config{}

	err := env.Parse(c)
	if err != nil {
		return nil, err
	}

	flag.IntVar(&c.QMax, "max", 6, "max number of task")
	flag.StringVar(&c.cFile, "p", "", "path to configuration file")
	flag.StringVar(&c.ServerAddr, "a", ":8080", "server address")

	flag.Parse()

	if c.cFile != "" {
		f, err := os.ReadFile(c.cFile)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(f, &c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

type Config struct {
	cFile string `env:"PATH"` // config file

	ServerAddr string `env:"ADDRESS" json:"server_addr,omitempty"`
	QMax       int    `env:"QMAX" json:"q_max,omitempty"`
}
