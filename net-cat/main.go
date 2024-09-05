package main

import (
	"fmt"
	"os"

	server "net_cat/server"
)

func checkValidPort(port string) int {
	if len(port) == 0 {
		return 1
	}
	for _, c := range port {
		if c < '0' || c > '9' {
			return 1
		}
	}
	return 0
}

func main() {
	args := os.Args[1:]
	port := "8989"

	if len(args) > 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(1)
	}
	if len(args) == 1 {
		if checkValidPort(args[0]) == 1 {
			fmt.Println("Port is invalid\n[USAGE]: ./TCPChat $port")
			os.Exit(1)
		}
		port = args[0]
	}
	server.Listen(port)
}
