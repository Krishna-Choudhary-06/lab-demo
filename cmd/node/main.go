package main

import (
	"fmt"
	"lab-demo/internal/network"
)

func main() {

	node := &network.Node{
		ID:      "node-B",
		Address: ":5001",
	}

	fmt.Println("Starting Node B...")

	err := node.Start()
	node.Connect("10.120.226.84:5000")

	if err != nil {
		fmt.Println("Node failed:", err)
		return
	}

	select {}
}
