package handlers

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

func EchoHandler(conn net.Conn) {
	defer func () {
		conn.Close()
		log.Printf("Connection from %s closed", conn.RemoteAddr())
	}()

	log.Printf("Received connection from %s", conn.RemoteAddr())
	buf := make([]byte, 8192)
	bytesRead, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Received %d bytes from %s:", bytesRead, conn.RemoteAddr())
	fmt.Println()
	fmt.Println(hex.Dump(buf[:bytesRead]))
	if bytesRead == len(buf) {
		log.Printf("Note: Any bytes greater than %d are ignored", len(buf))
	}

	bytesWritten, err := conn.Write(buf[:bytesRead])
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Echoed %d bytes back to %s:", bytesWritten, conn.RemoteAddr())
}