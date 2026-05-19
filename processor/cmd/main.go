package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/femitubosun/streaming-pipeline/processor/internal/config"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Could not load env")
		os.Exit(1)
	}
	fmt.Println("KAFKA_BROKER: ", cfg.KafkaBroker)

	// Wait for interrupt or termination signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("Shutting down...")

}
