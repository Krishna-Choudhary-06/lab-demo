package network

import (
	"fmt"
	"net"
)

func startServer(address string) (net.Listener, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	fmt.Println("TCP server listening on", address)

	return listener, nil
}
