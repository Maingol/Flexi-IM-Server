package main

import (
	"flag"
	"github.com/Maingol/Flexi-IM-Server/pkg/server"
)

func main() {
	tcpAddress := flag.String("tcp", "localhost:3000", "TCP server address")
	wsAddress := flag.String("ws", "localhost:4000", "WebSocket server address")
	flag.Parse()

	go server.StartTCPServer(*tcpAddress)
	server.StartWebSocketServer(*wsAddress)
}
