package config

import "math/big"

type Config struct {
	LogFile string
	Debug   bool

	Server struct {
		Port uint
	}
	DB struct {
		Path string
	}
	RSK struct {
		Endpoint string
		ChainId  *big.Int
	}
	Boltz struct {
		Endpoint string
	}
	Accounts struct {
		RSK struct {
			PrivateKey string
			Address    string
		}
	}
}
