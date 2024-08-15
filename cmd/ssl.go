package cmd

import (
	"crypto/tls"
	"fmt"
	"os"
	"log"
	"crypto/x509"

	"github.com/spf13/cobra"
)

var (
	sslCmd = &cobra.Command{
		Use:   "ssl",
		Short: "Run a SSL Socket Server (echo server)",
		Run: func(cmd *cobra.Command, args []string) {
			ip, err := cmd.Flags().GetString("ip")
			if err != nil {
				log.Fatal(err)
			}

			port, err := cmd.Flags().GetInt("port")
			if err != nil {
				log.Fatal(err)
			}
			cert := CheckPathFlagAndEnsureExists(cmd, "cert")
			key := CheckPathFlagAndEnsureExists(cmd, "key")
			ca := CheckPathFlagAndEnsureExists(cmd, "ca")
			RunSslServer(ip, port, cert, key, ca)
		},
	}
)

func RunSslServer(ip string, port int, cert string, key string, ca string) {
	//
	// Parse the cert and key pair
	//
	pair, err := tls.LoadX509KeyPair(cert, key)
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

	ca_pem, err := os.ReadFile(ca)
	if err != nil {
		log.Printf("Failed to read root CA file: %s", ca)
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

	address := fmt.Sprintf("%s:%d", ip, port)
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

func init() {
	sslCmd.Flags().String("ip", "127.0.0.1", "Listen IP")
	sslCmd.Flags().Int("port", 6666, "Listen Port")
	sslCmd.Flags().String("cert", "", "SSL Certificate")
	sslCmd.Flags().String("key", "", "SSL Key")
	sslCmd.Flags().String("ca", "", "CA Certificate")

	rootCmd.AddCommand(sslCmd)
}
