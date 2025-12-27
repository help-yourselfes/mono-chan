package config

import "time"

type Config struct {
	Env         string `yaml:"env" env-default:"development"`
	StoragePath string `yaml:"storage_path" env-default:"./storage/storage.db"`
	HTTPServer  `yaml:"server"`
	Services    struct {
		Gateway
		Boards
	}
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" end-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Gateway struct {
	URL         string        `yaml:"url" env-default:"localhost:3000"`
	Timeout     time.Duration `yaml:"timeout" end-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Boards struct {
	URL string `yaml:"url" env-default:"http://board-service:8081"`
}
