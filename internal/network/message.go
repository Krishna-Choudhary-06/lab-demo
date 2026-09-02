package network

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

type Message struct {
	Type     string `json:"type"`
	SenderID string `json:"sender_id"`
	Payload  string `json:"payload"`
}

func SendMessage(conn net.Conn, message Message) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	data = append(data, '\n')

	_, err = conn.Write(data)
	return err
}

func ReceiveMessages(conn net.Conn, handler func(Message)) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		var message Message

		err := json.Unmarshal(scanner.Bytes(), &message)
		if err != nil {
			fmt.Println("Invalid message:", err)
			continue
		}

		handler(message)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Connection read error:", err)
	}
}
