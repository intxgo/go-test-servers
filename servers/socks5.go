package servers

import (
	"fmt"
	"log"
	"net"

	"go-test-servers/config"
	"tailscale.com/net/socks5"
)

func RunSocksServer(config config.ServerConfig, status chan bool) {
	// Create a SOCKS5 server
	server := &socks5.Server{}

	// Enable authentication
	if config.Username != "" && config.Password == "" {
		log.Println("Password is required if username is provided")
		status <- false
		return
	}

	if config.Username == "" && config.Password != "" {
		log.Println("Username is required if password is provided")
		status <- false
		return
	}

	if config.Username != "" && config.Password != "" {
		log.Printf("Enabling authentication")
		log.Printf("  username: %s", config.Username)
		log.Printf("  password: %s", config.Password)
		server.Username = config.Username
		server.Password = config.Password
	}

	// Start a Listener
	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Println(err)
		status <- false
		return
	}

	// Create SOCKS5 proxy on localhost port 8000
	log.Printf("%s Listening on %s://%s", config.Type, config.Protocol, address)
	status <- true
	if err := server.Serve(l); err != nil {
		log.Fatal(err)
	}
}