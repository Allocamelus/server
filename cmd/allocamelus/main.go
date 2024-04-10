package main

import (
	"flag"
	"os"
	"strings"

	"github.com/allocamelus/server"
)

func main() {
	var (
		address string
	)

	flag.StringVar(&address, "a", "", "The address to listen on.")
	flag.Parse()

	if address != "" {
		os.Setenv("ADDRESS", strings.TrimSpace(address))
	}

	config := &server.Config{
		Address: strings.TrimSpace(os.Getenv("ADDRESS")),
	}
	config.Default()

	server.Start(config)
}
