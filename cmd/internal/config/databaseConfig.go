package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Circumference struct {
	CenterCoordinateX float64 `yaml:"center_сoordinate_x"`
	CenterCoordinateY float64 `yaml:"center_сoordinate_y"`
	Diameter          float64 `yaml:"diameter"`
}

func GetDatabaseConfig() *Circumference {

	cfg := &Circumference{}
	if err := cleanenv.ReadConfig("config.yml", cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
