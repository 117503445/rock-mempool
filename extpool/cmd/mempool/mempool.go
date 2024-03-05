package main

import (
	"log"
	"os"
	"rock-chain/extpool/internal/server"

	"github.com/caarlos0/env/v10"
	"github.com/rs/zerolog"
)

type config struct {
	Host   string   `env:"HOST" envDefault:"localhost:50051"`
	Others []string `env:"OTHERS" envSeparator:","`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
	}

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	serverCfg := &server.Config{
		Host:   cfg.Host,
		Others: cfg.Others,
		Logger: &logger,
	}

	server.NewServer(serverCfg).Run()
}
