package servers

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"go-test-servers/config"
	"go-test-servers/servers/handlers"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func RunSslSocketServer(cfg config.ServerConfig, status chan bool) {
	//
	// Parse the cert and key pair
	//
	if !FileExists(cfg.Cert) {
		log.Printf("Cannot start ssl server, Cert file not found: %s", cfg.Cert)
		status <- false
		return
	}

	if !FileExists(cfg.Key) {
		log.Printf("Cannot start ssl server, Key file not found: %s", cfg.Key)
		status <- false
		return
	}

	pair, err := tls.LoadX509KeyPair(cfg.Cert, cfg.Key)
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

	if !FileExists(cfg.Ca) {
		log.Printf("Cannot start ssl server, CA file not found: %s", cfg.Ca)
		status <- false
		return
	}

	if cfg.Ca != "" {
		log.Printf("Attempting to add CA Cert: %s", cfg.Ca)
		ca_pem, err := os.ReadFile(cfg.Ca)
		if err != nil {
			log.Printf("Failed to read root CA file: %s", cfg.Ca)
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

	var minTlsVersion uint16 = tls.VersionTLS10
	if cfg.MinTlsVersion != "" {
		minTlsVersion, err = config.ParseTlsVersion(cfg.MinTlsVersion)
		if err != nil {
			log.Printf("Failed to parse min TLS version: %s", cfg.MinTlsVersion)
			status <- false
			return
		}
	}
	var maxTlsVersion uint16 = tls.VersionTLS13
	if cfg.MaxTlsVersion != "" {
		maxTlsVersion, err = config.ParseTlsVersion(cfg.MaxTlsVersion)
		if err != nil {
			log.Printf("Failed to parse max TLS version: %s", cfg.MaxTlsVersion)
			status <- false
			return
		}
	}
	if minTlsVersion > maxTlsVersion {
		log.Printf("Min TLS version cannot be greater than max TLS version")
		status <- false
		return
	}

	cipherSuites := make([]uint16, len(cfg.CipherSuites))
	if cfg.CipherSuites != nil {
		for i, cipher := range cfg.CipherSuites {
			cipherSuites[i], err = config.ParseCipherSuite(cipher)
			if err != nil {
				log.Printf("Failed to parse cipher suite: %s", cipher)
				status <- false
				return
			}
		}
	}

	tlsConfig := &tls.Config{
		RootCAs:      rootCAs,
		Certificates: []tls.Certificate{pair},
		MinVersion:   minTlsVersion,
		MaxVersion:   maxTlsVersion,
		CipherSuites: cipherSuites,
	}

	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	listener, err := tls.Listen("tcp", address, tlsConfig)
	if err != nil {
		log.Println(err)
		status <- false
		return
	}

	// Determine Handler Type
	var connHandler handlers.ConnectionHandler
	switch cfg.HandlerType {
	case config.Echo:
		connHandler = handlers.EchoHandler
	default:
		log.Printf("Unknown handler type %s, using echo handler", cfg.HandlerType)
		connHandler = handlers.EchoHandler
	}

	log.Printf("%s Listening on %s", cfg.Type, address)
	status <- true
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go connHandler(conn)
	}
}
