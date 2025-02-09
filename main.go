package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: chaincode <command>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "suggest":
		fmt.Println("[AI]: Generating Golang or Solidity code snippet...")
	case "debug":
		fmt.Println("[AI]: Debugging your Solidity code...")
	default:
		fmt.Println("Unknown command. Available commands: suggest, debug")
	}
}
