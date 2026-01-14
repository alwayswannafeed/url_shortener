package main

import (
	"os"

	"github.com/alwayswannafeed/url_shortener/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}