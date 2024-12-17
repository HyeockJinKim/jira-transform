package config

import (
	"io"
	"os"

	"github.com/BurntSushi/toml"
)

const (
	envConfigPath = "CONFIG_PATH"

	defaultConfigPath = "configs/config.toml"
)

type Config struct {
	HTTP   HTTP   `toml:"http"`
	Mirror Mirror `toml:"mirror"`
}

type HTTP struct {
	Addr              string `toml:"addr"`
	ShutdownTimeoutMs int64  `toml:"shutdownTimeoutMs"`
}

type Mirror struct {
	Github Github `toml:"github"`
}

type Github struct {
	Owner string `toml:"owner"`
	Token string `toml:"token"`
}

func ReadConfig() (*Config, error) {
	var cfg Config
	path, prs := os.LookupEnv(envConfigPath)
	if !prs {
		path = defaultConfigPath
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	if err := toml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
