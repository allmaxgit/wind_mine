package configs

import (
	"github.com/BurntSushi/toml"
)

type Configs struct {
	Common   Common
	Server   Server
	Crypto   Crypto
	Services Services
	DB       map[string]DB `toml:"database"`
}

type Common struct {
	Dev         bool
	LogsOutPath string
}

type Server struct {
	TCPPort uint `toml:"tcpPort"`
}

type Crypto struct {
	BTCAddr string `toml:"btcAddress"`
}

type Services struct {
	BTCServicePort uint `toml:"btcServicePort"`
}

type DB struct {
	Host         string
	Port         uint
	Name         string
	User         string
	Password     string
}

var configs Configs

// ParseConfigs parses and returns configs from toml file.
func ParseConfigs(confPath string) (*Configs, error) {
	if _, err := toml.DecodeFile(confPath, &configs); err != nil {
		return nil, err
	}

	return &configs, nil
}

// GetConfigs return configs.
// You can get it from any place
// func GetConfigs() *Configs { return &configs }
