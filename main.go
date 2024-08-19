package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"go-test-servers/config"
	"go-test-servers/servers"

	"gopkg.in/yaml.v3"
)

type StartServerFunc func(config.ServerConfig, chan bool)

func StartServer(config config.ServerConfig, serverFunc StartServerFunc) {
	status := make(chan bool)
	fmt.Println()
	log.Printf("Attempting to start %s server\n", config.Type)
	go serverFunc(config, status)
	if <-status {
		log.Printf("Started %s server\n", config.Type)
	} else {
		log.Printf("Failed to start %s server\n", config.Type)
	}
	fmt.Println()
}

func ReadConfig(configPath *string) config.Config {
	if *configPath == "" {
		log.Fatal("No config file specified")
	}

	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		log.Fatalf("Config file not found: %s", *configPath)
	}

	yamlFile, err := os.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("Failed to read config file: %s, err: %v", *configPath, err)
	}

	serverConfig := config.Config{}
	err = yaml.Unmarshal(yamlFile, &serverConfig)
	if err != nil {
		log.Fatalf("Failed to unmarshal config file: %s, err: %v", *configPath, err)
	}

	return serverConfig
}

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

	configPath := flag.String("config", "config.yaml", "path to the config file")
	flag.Parse()

	serverConfig := ReadConfig(configPath)
	// Start the servers
	if len(serverConfig.Servers) == 0 {
		log.Println("No servers defined in the config")
		os.Exit(1)
	}

	for _, server := range serverConfig.Servers {
		if !server.Enabled {
			log.Printf("Server is disabled in the config, skipping:  %s\n", server.Type)
			continue
		}
		switch server.Type {
		case config.Socket:
			StartServer(server, servers.RunTcpSocketServer)
		case config.Socks5:
			StartServer(server, servers.RunSocksServer)
		case config.Ssl:
			StartServer(server, servers.RunSslSocketServer)
		default:
			log.Printf("Unknown server type: %s\n", server.Type)
		}
	}

	//wait forever
	log.Println("Waiting, press Ctrl+C to exit")
	select {}
}
