package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type ServerConfig struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func GetServerConfig() *ServerConfig {
	cfg := &ServerConfig{}
	if err := cleanenv.ReadConfig("config.yml", cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
