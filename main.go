package main

import (
	"flag"
	"log"
	"net"
	"tailscale.com/net/socks5"
)

func main() {
	user := flag.String("user", "", "Username")
	pass := flag.String("pass", "", "Password")
	network := flag.String("network", "tcp", "Network type")
	addr := flag.String("addr", "127.0.0.1:8000", "Address to listen on")

	flag.Parse()

	if *user == "" && *pass != "" || *user != "" && *pass == "" {
		log.Fatalf("Both username and password must be provided to enable authentication")
	}

	server := &socks5.Server{}

	if *user != "" && *pass != "" {
		log.Printf("Enabling authentication")
		log.Printf("  Username: %s", *user)
		log.Printf("  Password: %s", *pass)
		server.Username = *user
		server.Password = *pass
	}


	l, err := net.Listen(*network, *addr)
	if err != nil {
		log.Fatal(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	log.Printf("Serving on %s://%s", *network, *addr)
	if err := server.Serve(l); err != nil {
		log.Fatal(err)
	}

}