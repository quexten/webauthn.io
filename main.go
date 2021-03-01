package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/quexten/webauthn.io/config"
	log "github.com/quexten/webauthn.io/logger"
	"github.com/quexten/webauthn.io/models"
	"github.com/quexten/webauthn.io/server"
)

func main() {
	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = models.Setup(config)
	if err != nil {
		log.Fatal(err)
	}

	err = log.Setup(config)
	if err != nil {
		log.Fatal(err)
	}

	server, err := server.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}
	go server.Start()

	// Handle graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGINT)

	<-c
	log.Info("Shutting down...")
	server.Shutdown()
}
