package servers

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"go-test-servers/config"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func RunSslSocketServer(config config.ServerConfig, status chan bool) {
	//
	// Parse the cert and key pair
	//
	if !FileExists(config.Cert) {
		log.Printf("Cannot start ssl server, Cert file not found: %s", config.Cert)
		status <- false
		return
	}

	if !FileExists(config.Key) {
		log.Printf("Cannot start ssl server, Key file not found: %s", config.Key)
		status <- false
		return
	}

	pair, err := tls.LoadX509KeyPair(config.Cert, config.Key)
	if err != nil {
		log.Printf("Failed to load server cert/key pair")
		status <- false
		return
	}

	//
	// Load CA Certs
	//
	log.Printf("Loading system CA Certs")
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	if !FileExists(config.Ca) {
		log.Printf("Cannot start ssl server, CA file not found: %s", config.Ca)
		status <- false
		return
	}

	if (config.Ca != "") {
		log.Printf("Attempting to add CA Cert: %s", config.Ca)
		ca_pem, err := os.ReadFile(config.Ca)
		if err != nil {
			log.Printf("Failed to read root CA file: %s", config.Ca)
			status <- false
			return
		}
		ok := rootCAs.AppendCertsFromPEM(ca_pem)
		if !ok {
			log.Printf("Failed to parse root CA certificate")
			status <- false
			return
		}
	}

	tlsConfig := &tls.Config{
		RootCAs : rootCAs,
		Certificates: []tls.Certificate{pair},
	}

	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	listener, err := tls.Listen("tcp", address, tlsConfig)
	if err != nil {
		log.Println(err)
		status <- false
		return
	}

	log.Printf("%s Listening on %s", config.Type, address)
	status <- true
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go HandleSocketConnection(conn)

	}
}