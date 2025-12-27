package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

func Load() *Config {
	cfgPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *cfgPath == "" {
		if v, ok := os.LookupEnv("CONFIG_PATH"); ok && v != "" {
			*cfgPath = v
		}
	}

	if *cfgPath == "" {
		*cfgPath = "./config/config.yaml"
	}

	if _, err := os.Stat(*cfgPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", *cfgPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(*cfgPath, &cfg); err != nil {
		log.Fatalf("Cannot read config: %s", err)
	}

	return &cfg
}
