package main

import (
	"go-test-servers/cmd"
	"log"
	"os"
	"os/signal"
)

func main() {
	// Handle SIGINT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("captured %v, stopping server and exiting..", sig)
			os.Exit(1)
		}
	}()

	cmd.Execute()
}
