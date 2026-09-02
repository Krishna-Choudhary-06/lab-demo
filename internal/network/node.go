package network

import (
	"fmt"
	"net"
)

type Node struct {
	ID           string
	Address      string
	OutgoingConn net.Conn
}

func (n *Node) Start() error {
	listener, err := startServer(n.Address)
	if err != nil {
		return err
	}

	fmt.Println("Node", n.ID, "started")

	go func() {
		defer listener.Close()

		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept error:", err)
				return
			}

			go n.handleConnection(conn)
		}
	}()

	return nil
}

func (n *Node) Connect(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	fmt.Println("Node", n.ID, "connected to", address)

	n.OutgoingConn = conn

	return nil
}

func (n *Node) SendMessage(conn net.Conn, message string) error {
	_, err := conn.Write([]byte(message))
	return err
}

func (n *Node) handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		bytesRead, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("Peer disconnected:", conn.RemoteAddr())
			return
		}

		message := string(buffer[:bytesRead])

		fmt.Println("Message received:", message)
	}
}
