package servers

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"go-test-servers/config"
)

func RunSslSocketServer(config config.SslSocketConfig) {
	//
	// Parse the cert and key pair
	//
	pair, err := tls.LoadX509KeyPair(config.Cert, config.Key)
	if err != nil {
		log.Printf("Failed to load server cert/key pair")
		log.Fatal(err)
	}

	//
	// Load CA Certs
	//
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	ca_pem, err := os.ReadFile(config.Ca)
	if err != nil {
		log.Printf("Failed to read root CA file: %s", config.Ca)
		log.Fatal(err)
	}
	ok := rootCAs.AppendCertsFromPEM(ca_pem)
	if !ok {
		log.Fatal("Failed to parse root CA certificate")
	}

	tlsConfig := &tls.Config{
		RootCAs : rootCAs,
		Certificates: []tls.Certificate{pair},
	}

	address := fmt.Sprintf("%s:%d", config.Host, config.Port)
	listener, err := tls.Listen("tcp", address, tlsConfig)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Ssl Server Listening on %s", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go HandleSocketConnection(conn)

	}
}