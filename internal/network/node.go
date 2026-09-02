package network

import (
	"fmt"
	"net"
)

type Node struct {
	ID          string
	Address     string
	PeerManager *PeerManager
}

func (n *Node) Start() error {
	listener, err := startServer(n.Address)
	if err != nil {
		return err
	}

	n.PeerManager = NewPeerManager()

	fmt.Println("Node", n.ID, "started")

	go func() {
		defer listener.Close()

		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept error:", err)
				continue
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

	peer := &Peer{
		Address:   address,
		Conn:      conn,
		Connected: true,
	}

	n.PeerManager.AddPeer(peer)

	fmt.Println("Node", n.ID, "connected to", address)

	return nil
}

func (n *Node) SendMessage(conn net.Conn, message Message) error {
	return SendMessage(conn, message)
}

func (n *Node) handleConnection(conn net.Conn) {
	peer := &Peer{
		Address:   conn.RemoteAddr().String(),
		Conn:      conn,
		Connected: true,
	}

	n.PeerManager.AddPeer(peer)

	fmt.Println("Incoming connection from:", peer.Address)

	ReceiveMessages(conn, func(message Message) {
		fmt.Println("Message received:")
		fmt.Println("  Type:", message.Type)
		fmt.Println("  Sender:", message.SenderID)
		fmt.Println("  Payload:", message.Payload)
	})

	fmt.Println("Peer disconnected:", peer.Address)

	n.PeerManager.RemovePeer(peer.Address)

	conn.Close()
}
