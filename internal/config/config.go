package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	ENV        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8000"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("config path must not by empty")
	}

	if _, err := os.Stat(configPath); os.IsExist(err) {
		log.Fatalf("file %s dose not exist", configPath)
	}

	var cnf Config

	if err := cleanenv.ReadConfig(configPath, &cnf); err != nil {
		log.Fatalf("can not read config %s", err)
	}
	return cnf
}
