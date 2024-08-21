package servers

import (
	"fmt"
	"log"
	"net"

	"go-test-servers/config"
	"go-test-servers/servers/handlers"
)

func RunTcpSocketServer(cfg config.ServerConfig, status chan bool) {
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	listener, err := net.Listen("tcp", address)
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
		connHandler= handlers.EchoHandler
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