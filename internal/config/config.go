package config

import "github.com/joho/godotenv"

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}

type GRPCConfig interface {
	Address() string
}

type HTTPConfig interface {
	Address() string
}

type DBConfig interface {
	DSN() string
}
