package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
)

var (
	tcpCmd = &cobra.Command{
		Use:   "tcp",
		Short: "Run a TCP Socket Server (echo server)",
		Run: func(cmd *cobra.Command, args []string) {
			ip, err := cmd.Flags().GetString("ip")
			if err != nil {
				log.Fatal(err)
			}

			port, err := cmd.Flags().GetInt("port")
			if err != nil {
				log.Fatal(err)
			}
			RunTcpServer(ip, port)
		},
	}
)

func RunTcpServer(ip string, port int) {
	address := fmt.Sprintf("%s:%d", ip, port)	
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Tcp Server Listening on %s", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go HandleSocketConnection(conn)
	}
}

func init() {
	tcpCmd.Flags().String("ip", "127.0.0.1", "Listen IP")
	tcpCmd.Flags().Int("port", 7777, "Listen Port")
	rootCmd.AddCommand(tcpCmd)
}
