package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"tailscale.com/net/socks5"
)

var (
	socks5Cmd = &cobra.Command{
		Use:   "socks5",
		Short: "Run a SOCKS5 Server",
		Run: func(cmd *cobra.Command, args []string) {
			user, err := cmd.Flags().GetString("user")
			if err != nil {
				log.Fatal(err)
			}

			pass, err := cmd.Flags().GetString("pass")
			if err != nil {
				log.Fatal(err)
			}

			proto, err := cmd.Flags().GetString("proto")
			if err != nil {
				log.Fatal(err)
			}

			ip, err := cmd.Flags().GetString("ip")
			if err != nil {
				log.Fatal(err)
			}

			port, err := cmd.Flags().GetInt("port"); if err != nil {
				log.Fatal(err)
			}
			RunSocksServer(user, pass, proto, ip, port)
		},
	}
)

func RunSocksServer(user string, pass string, proto string, ip string, port int) {
	// Create a SOCKS5 server
	server := &socks5.Server{}

	// Enable authentication
	if user != "" && pass == "" {
		log.Fatal("Password is required if username is provided")
	}
	if user == "" && pass != "" {
		log.Fatal("Username is required if password is provided")
	}
	if user != "" && pass != "" {
		log.Printf("Enabling authentication")
		log.Printf("  Username: %s", user)
		log.Printf("  Password: %s", pass)
		server.Username = user
		server.Password = pass
	}

	// Start a Listener
	address := fmt.Sprintf("%s:%d", ip, port)
	l, err := net.Listen(proto, address)
	if err != nil {
		log.Fatal(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	log.Printf("SOCKS5 Listening on %s://%s", proto, address)
	if err := server.Serve(l); err != nil {
		log.Fatal(err)
	}

}

func init() {
	socks5Cmd.Flags().String("user", "", "Username")
	socks5Cmd.Flags().String("pass", "", "Password")
	socks5Cmd.Flags().String("proto", "tcp", "Protocol")
	socks5Cmd.Flags().String("ip", "127.0.0.1", "Listen IP")
	socks5Cmd.Flags().Int("port", 8888, "Listen Port")
	rootCmd.AddCommand(socks5Cmd)
}
