package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"go-test-servers/config"
	"go-test-servers/servers"

	"gopkg.in/yaml.v3"
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

	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	serverConfig := config.Config{}
	err = yaml.Unmarshal(yamlFile, &serverConfig)
	if err != nil {
		panic(err)
	}

	// Start the servers
	log.Println("Server Count : ", len(serverConfig.Servers))
	for _, server := range serverConfig.Servers {
		if !server.Enabled {
			log.Printf("Server is disabled in the config, skipping:  %s\n", server.Type)
			continue
		}

		fmt.Println() // place a new line between each server's startup for readability
		status := make(chan bool)
		log.Printf("Attempting to start %s\n", server.Type)
		switch server.Type {
		case config.Socket:
			go servers.RunTcpSocketServer(server, status)
		case config.Socks5:
			go servers.RunSocksServer(server, status)
		case config.Ssl:
			go servers.RunSslSocketServer(server, status)
		default:
			log.Printf("Unknown server type: %s\n", server.Type)
		}

		if <-status {
			log.Printf("Started %s server\n", server.Type)
		} else {
			log.Printf("Failed to start %s server\n", server.Type)
		}
		fmt.Println()
	}

	//wait forever
	log.Println("Waiting, press Ctrl+C to exit")
	select {}
}
