package main

import (
	"fmt"
	"lab-demo/internal/network"
)

func main() {

	node := &network.Node{
		ID:      "node-A",
		Address: ":5000",
	}

	fmt.Println("Starting Node A...")

	err := node.Start()
	node.Connect("10.229.64.202:5001")

	if err != nil {
		fmt.Println("Node failed:", err)
		return
	}

	select {}
}
