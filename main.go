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

	config := config.Config{}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	if config.Socks5.Enabled {
		fmt.Println("Starting SOCKS5 Server")
		go servers.RunSocksServer(config.Socks5)
	}

	if config.TcpSocket.Enabled {
		fmt.Println("Starting Socket Server")
		go servers.RunTcpSocketServer(config.TcpSocket)
	}

	if config.SslSocket.Enabled {
		fmt.Println("Starting SSL Server")
		go servers.RunSslSocketServer(config.SslSocket)
	}

	//wait forever
	fmt.Println("Waiting forever")
	select {}
}
