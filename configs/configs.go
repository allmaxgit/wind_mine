package configs

import (
	"github.com/BurntSushi/toml"
	"math/big"
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
	// BTC
	BTCAddr          string `toml:"btcAddress"`

	// ETH
	ETHNetworkId     int `toml:"ethNetworkId"`
	InfuraToken      string

	// Contract
	CrowdsaleAddress string
	OwnerPrivateKey  string
	OwnerAddress     string
	GasPrice         big.Int
	GasLimit         uint64
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
