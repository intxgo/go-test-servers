package handlers

import (
	"net"
)

type ConnectionHandler func(net.Conn)
