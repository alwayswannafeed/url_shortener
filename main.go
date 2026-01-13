package main

import (
	"os"

	"github.com/alwayswannafeed/url_shortener/internal/config"
	"github.com/alwayswannafeed/url_shortener/internal/service"
	"gitlab.com/distributed_lab/kit/kv"
)

func main() {
	if len(os.Args) == 1 {
		os.Setenv("KV_VIPER_FILE", "config.yaml")
	}

	cfg := config.New(kv.MustFromEnv())

	service.Run(cfg)
}